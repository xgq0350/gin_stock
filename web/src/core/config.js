import { ref } from 'vue'

/**
 * 网站配置文件
 */

const config = {
  appName: 'Gin-Vue-Admin',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Gin-Vue-Admin，开源地址：https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:v2.5.6`
      )
    )
    console.log(
      chalk.green(
        `> 加群方式:微信：shouzi_1994 QQ群：622360840`
      )
    )
    console.log(
      chalk.green(
        `> GVA讨论社区：https://support.qq.com/products/371961`
      )
    )
    console.log(
      chalk.green(
        `> 插件市场:https://plugin.gin-vue-admin.com`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log(
      chalk.green(
        `> 如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/coffee/index.html`
      )
    )
    console.log('\n')
  }
}

export const cwfxOptions = ref([
  {
    value: '三线蓝筹',
    label: '三线蓝筹',
  },
  {
    value: '二线蓝筹',
    label: '二线蓝筹',
  },
  {
    value: '权重股',
    label: '权重股',
  },
  {
    value: '破净股',
    label: '破净股',
  },
  {
    value: '白马股',
    label: '白马股',
  },
  {
    value: '绩优股',
    label: '绩优股',
  },
  {
    value: '预戴帽',
    label: '预戴帽',
  },
  {
    value: '绩差股',
    label: '绩差股',
  }
])

export const close2023 = (data) => {
  return data.close_max_2023 + ' / ' + data.close_min_2023
}

export const close2022 = (data) => {
  return data.close_max_2022 + ' / ' + data.close_min_2022
}

export const close2021 = (data) => {
  return data.close_max_2021 + ' / ' + data.close_min_2021
}

export const getParams = () => {
  var url = window.location.href;
  const paramsObj = ref({})
  const params = url.split('?')[1];
  if (params) {
    let paramsArr = params.split('&');
    for (let i = 0; i < paramsArr.length; i++) {
      let param = paramsArr[i].split('=');
      paramsObj[param[0]] = param[1];
    }
    return paramsObj
  }
 return ''
}

export default config
