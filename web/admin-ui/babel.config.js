module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],
  "plugins": [
    [
      "component",
      {
        "libraryName": "element-ui",
        "styleLibraryName": "theme-chalk"
      }
    ]
  ]
  // "plugins": [["import", {
  //   "libraryName": "view-design",
  //   "libraryDirectory": "src/components"
  // }]]
}
