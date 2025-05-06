require 'json'
require 'fileutils'

# コマンドライン引数からファイルパスを取得
if ARGV.length < 1
  puts "使用方法: ruby create_test_data.rb <出力ディレクトリ> [データ数=5000]"
  puts "例: ruby create_test_data.rb ../test_cases/case6 10000"
  exit 1
end

output_dir = ARGV[0]
data_count = ARGV[1] ? ARGV[1].to_i : 5000

# 出力ディレクトリが存在しない場合は作成
FileUtils.mkdir_p(output_dir) unless Dir.exist?(output_dir)

# input.txt用のデータ生成
operations = []

# 基本的なデータを追加
operations << { "action" => "put", "key" => "apple", "value" => 5 }
operations << { "action" => "put", "key" => "banana", "value" => 8 }
operations << { "action" => "put", "key" => "orange", "value" => 10 }

# 大量のデータを生成
(1..data_count).each do |i|
  operations << { "action" => "put", "key" => "key#{i}", "value" => i * 100 }
end

# いくつかのgetとremove操作を追加
operations << { "action" => "get", "key" => "key1" }
operations << { "action" => "get", "key" => "key#{data_count/2}" }
operations << { "action" => "get", "key" => "key#{data_count}" }

# 一部のキーを削除
(1..10).each do |i|
  remove_key = rand(1..data_count)
  operations << { "action" => "remove", "key" => "key#{remove_key}" }
end

# expected.txt用のデータ生成
expected = {
  "apple" => 5,
  "banana" => 8,
  "orange" => 10
}

# 残りのキーと値を追加
(1..data_count).each do |i|
  expected["key#{i}"] = i * 100
end

# 削除したキーを除外（最後の10個のremove操作）
operations.last(10).each do |op|
  if op["action"] == "remove"
    expected.delete(op["key"])
  end
end

# input.txtに書き込み
File.open(File.join(output_dir, "input.txt"), "w") do |f|
  f.puts JSON.pretty_generate(operations)
end

# expected.txtに書き込み
File.open(File.join(output_dir, "expected.txt"), "w") do |f|
  f.puts JSON.pretty_generate(expected)
end

puts "テストケースを生成しました:"
puts "- 入力ファイル: #{File.join(output_dir, "input.txt")}"
puts "- 期待出力ファイル: #{File.join(output_dir, "expected.txt")}"
puts "- 総データ数: #{data_count + 3} エントリ (基本データ3 + 生成データ#{data_count})"
puts "- 削除されたエントリ: 10"
puts "- 最終的なエントリ数: #{expected.size}"
