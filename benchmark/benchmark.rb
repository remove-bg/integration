# frozen_string_literal: true

# Usage:
# REMOVE_BG_API_KEY=your-api-key ruby benchmark.rb

require "bundler/inline"
require "benchmark"

gemfile do
  gem "remove_bg"
end

api_key = ENV.fetch("REMOVE_BG_API_KEY")
iterations = ENV.fetch("ITERATIONS", 10).to_i
image_url = ENV.fetch("IMAGE_URL", "https://images.pexels.com/photos/1250643/pexels-photo-1250643.jpeg")
image_sizes = ENV["IMAGE_SIZES"]&.split(",") || ["regular", "medium", "hd", "4k"]

process_image = Proc.new do |image_size|
  RemoveBg.from_url(image_url, api_key: api_key, size: image_size)
end

GC.disable

results = Benchmark.bm do |benchmark|
  image_sizes.each do |image_size|
    benchmark.report(image_size) do
      iterations.times { process_image.call(image_size) }
    end
  end
end

puts "\nAverage duration (seconds) per image:"
results.each do |result|
  avg = (result.real / iterations).round(3)
  puts "#{result.label.ljust(10)} #{avg}"
end
