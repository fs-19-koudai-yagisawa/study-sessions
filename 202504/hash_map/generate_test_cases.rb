#!/usr/bin/env ruby

require 'json'
require 'fileutils'

# Function to generate test cases
def generate_test_case(case_number, operation_count)
  base_path = File.join(Dir.pwd, "hash_map", "test_cases", "case#{case_number}")
  FileUtils.mkdir_p(base_path) unless Dir.exist?(base_path)
  
  puts "Generating test case #{case_number} with #{operation_count} operations..."
  
  # Generate random keys and values
  keys = (0...operation_count).map { |i| "key#{i}" }
  
  # Create a mix of put, get, and remove operations
  operations = []
  final_state = {}
  
  operation_count.times do |i|
    # Distribute operations: 60% put, 30% get, 10% remove
    op_type = if i < operation_count * 0.6
                :put
              elsif i < operation_count * 0.9
                :get
              else
                :remove
              end
    
    case op_type
    when :put
      key = keys.sample
      value = rand(1..100)
      operations << { action: "put", key: key, value: value }
      final_state[key] = value
    when :get
      # Only get keys that exist
      key = final_state.keys.sample || keys.sample
      operations << { action: "get", key: key }
    when :remove
      # Only remove keys that exist
      if final_state.size > 0
        key = final_state.keys.sample
        operations << { action: "remove", key: key }
        final_state.delete(key)
      else
        # If no keys exist, do a put instead
        key = keys.sample
        value = rand(1..100)
        operations << { action: "put", key: key, value: value }
        final_state[key] = value
      end
    end
  end
  
  # Write the input file
  File.write(File.join(base_path, "input.txt"), JSON.pretty_generate(operations))
  
  # Write the expected file
  File.write(File.join(base_path, "expected.txt"), JSON.pretty_generate(final_state))
  
  puts "Generated test case #{case_number} with #{operation_count} operations"
end

# Generate test cases with increasing operation counts
generate_test_case(3, 1000)
generate_test_case(4, 10000)
generate_test_case(5, 100000)

puts "All test cases generated successfully!"
