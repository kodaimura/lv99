import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "standalone",
  /* config options here */
  experimental: {
    authInterrupts: true,
  },
  devIndicators: false,
};

export default nextConfig;
