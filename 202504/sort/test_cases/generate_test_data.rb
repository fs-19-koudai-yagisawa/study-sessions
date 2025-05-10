#!/usr/bin/env ruby

require 'json'
require 'fileutils'

# テストケース生成用のクラス
class TestCaseGenerator
  def initialize
    @base_dir = File.dirname(__FILE__)
  end

  # テストケースを生成する
  def generate_test_case(case_name, data, description = nil)
    case_dir = File.join(@base_dir, case_name)
    Dir.mkdir(case_dir) unless Dir.exist?(case_dir)

    # 説明を保存
    if description
      File.open(File.join(case_dir, "description.txt"), "w") do |f|
        f.puts description
      end
    end

    # 入力データを保存
    File.open(File.join(case_dir, "input.txt"), "w") do |f|
      f.puts "[#{data.join(", ")}]"
    end

    # 期待される結果（ソート済みデータ）を保存
    expected = data.sort
    File.open(File.join(case_dir, "expected.txt"), "w") do |f|
      f.puts "[#{expected.join(", ")}]"
    end
  end

  # すべてのテストケースを生成
  def generate_all
    generate_basic_cases
    generate_edge_cases
    generate_large_cases
    generate_special_cases
  end

  private

  # 基本的なテストケース
  def generate_basic_cases
    # 小さなデータセット
    generate_test_case("case1", [5, 1, 4, 2, 8],
      "基本的なテストケース：小さな正の整数の配列")

    # 負の数を含むデータ
    generate_test_case("case2", ["orange", "apple", "banana", "grape", "kiwi"],
      "基本的なテストケース：文字列の配列")

    # 重複を含むデータ
    generate_test_case("case3", [3.14, 1.41, 2.71, 1.73, 2.0],
      "基本的なテストケース：浮動小数点数の配列")
  end

  # エッジケース
  def generate_edge_cases
    # 空の配列
    generate_test_case("case_edge1", [],
      "エッジケース：空の配列")

    # 1要素の配列
    generate_test_case("case_edge2", [1],
      "エッジケース：1要素の配列")

    # すべての要素が同じ配列
    generate_test_case("case_edge3", [2, 2, 2, 2, 2],
      "エッジケース：すべての要素が同じ配列")
  end

  # 大規模なテストケース
  def generate_large_cases
    # 大規模な整数データ
    large_int_data = Array.new(10_000_000) { rand(-1_000_000_000..1_000_000_000) }
    generate_test_case("case_large1", large_int_data,
      "1000万件のランダムな整数（-10億から10億の範囲）")

    # 大規模な文字列データ
    chars = ('a'..'z').to_a
    large_str_data = Array.new(10_000_000) do
      word_length = rand(5..15)  # 5-15文字のランダムな長さ
      word_length.times.map { chars.sample }.join
    end
    generate_test_case("case_large2", large_str_data,
      "1000万件のランダムな文字列（5-15文字）")

    # TODO: 未実装 テストできない
    # 大規模な浮動小数点数データ
    # large_float_data = Array.new(10_000_000) do
    #   # -1000から1000の範囲で、小数点以下6桁までのランダムな浮動小数点数
    #   (rand(-1000.0..1000.0) * 1000000).round / 1000000.0
    # end
    # generate_test_case("case_large3", large_float_data,
    #   "1000万件のランダムな浮動小数点数（小数点以下6桁）")

    # すでにソート済みの大規模データ
    sorted_data = Array.new(10_000_000) { |i| i }
    generate_test_case("case_large4", sorted_data,
      "1000万件のソート済みデータ")

    # 逆順の大規模データ
    reverse_data = Array.new(10_000_000) { |i| 10_000_000 - i }
    generate_test_case("case_large5", reverse_data,
      "1000万件の逆順データ")
  end

  # 特殊なケース
  def generate_special_cases
    # 最大値と最小値を含むデータ
    generate_test_case("case_special1", [0, 1_000_000_000, -1_000_000_000, 5, -5],
      "特殊ケース：最大値と最小値を含む配列")

    # 浮動小数点数の特殊なケース
    generate_test_case("case_special2", [0.0, -0.0, 3.14, -3.14, 2.0e-10, 2.0e10],
      "特殊ケース：様々な浮動小数点数を含む配列")
  end
end

# テストケースを生成
generator = TestCaseGenerator.new
generator.generate_all
puts "テストケースの生成が完了しました。"
