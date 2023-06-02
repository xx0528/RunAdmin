/*
 * @Author: xx
 * @Date: 2023-04-24 10:53:04
 * @LastEditTime: 2023-05-09 19:44:01
 * @Description: 
 */
/**
 * 网站配置文件
 */

const config = {
  appName: 'RunAdmin',
  // appLogo: 'https://img.sj33.cn/uploads/allimg/201304/101F01D8-1.png',
  appLogo: '../assets/facebook.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Run-Admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:v2.5.5`
      )
    )
    console.log('\n')
  }
}

export default config
