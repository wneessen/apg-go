class ApgGo < Formula
  desc "\"Automated Password Generator\"-clone written in Go"
  homepage "https://github.com/wneessen/apg-go/"
  url "https://github.com/wneessen/apg-go/archive/refs/tags/v0.3.1.tar.gz"
  sha256 "1a798bd729c2985a11001118ad7d222e75c4f01e642184c37dec409a899a565b"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args
  end

  test do
    system "go", "test", "github.com/wneessen/apg-go"
  end
end
