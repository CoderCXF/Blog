module.exports = {
  transpileDependencies: true,
  publicPath: '/admin/',
  outputDir: 'dist',
  assetsDir: 'static',
  devServer: { // 环境配置
    host: '0.0.0.0',
    port: '8080',
    https: false,
    allowedHosts: 'all',
    open: false // 配置自动启动浏览器
  }
}
