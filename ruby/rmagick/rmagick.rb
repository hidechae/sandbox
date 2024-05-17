# frozen_string_literal: true
require 'rmagick'
require 'tempfile'

# 画像の読み込み
def load_image(file_path)
  Magick::Image.read(file_path).first
end

# カラープロファイルを残してExif情報を削除する
def strip_without_icc_profile(image)
  # 事前にカラープロファイルを取得する
  icc_profile = image.color_profile

  # Exif情報の削除
  image.strip!

  print icc_profile + "\n"

  # カラープロファイルを再設定
  image.color_profile = icc_profile

  # # カラープロファイルを再設定
  # if icc_profile
  #   # 拡張子が icc でなければ add_profile でエラーとなる
  #   Tempfile.open(%w[icc_profile .icc]) do |f|
  #     f.binmode
  #     f.write(icc_profile)
  #     f.flush
  #     image.add_profile(f.path)
  #   end
  # end
  image
end

# 画像の書き込み
def write_image(image, output_path)
  image.write(output_path)
end

if ARGV.empty? || ARGV.size < 2
  puts "Image path is required."
  puts "Usage: bundle exec ruby #{__FILE__} image_path output_path"
  exit
end

image_path = ARGV[0]
output_path = ARGV[1]
image = load_image(image_path)
image = strip_without_icc_profile(image)
write_image(image, output_path)
