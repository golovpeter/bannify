class Bannify < Formula
  desc "Generate colorful ASCII art banners in your terminal"
  homepage "https://github.com/golovpeter/bannify"
  url "https://github.com/golovpeter/bannify/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "8eb14c6240b9aec47505a2f730dd7a8d3c0e94304dd2ae36883a4ea8bef24b76"
  license "MIT"
  head "https://github.com/username/bannify.git", branch: "main"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w -X main.version=#{version}")
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/bannify -v")

    # Проверяем, что программа работает
    output = shell_output("#{bin}/bannify -t TEST 2>&1")
    assert_match "TEST", output
  end
end