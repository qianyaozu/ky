import axios from 'axios'
import Mock from 'mockjs'
import MockAdapter from 'axios-mock-adapter'
export default {
  bootstrap () {
    let mock = new MockAdapter(axios)
    mock.onPost('/api/login').reply(config => {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          resolve([200, {State: 1, Data: {UserName: 'admin', MineName: 'jeqee', Token: 123456789}, Message: ''}])
        }, 1000)
      })
    })

    mock.onGet('/api/getScripts').reply(config => {
      return new Promise((resolve, reject) => {
        var tableData = []
        for (let i = 0; i < 15; i++) {
          tableData.push({
            ID: Mock.mock('@string()'),
            Name: Mock.mock('@cname'),
            CreateTime: Mock.mock('@date("yyyy-MM-dd")')
          })
        }
        resolve([200, {State: 1, Data: tableData, Message: ''}])
      })
    })
  }
}
