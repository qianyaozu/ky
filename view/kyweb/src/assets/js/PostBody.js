export default function PostBody () {
  this.Method = ''  // get,count,update,add,delete
  this.DBName = ''
  this.Table = ''
  this.Data = ''
  this.OrderBy = ''
  this.Limit = 0
  this.Skip = 0
  this.Select = {}
  this.UpdateAll = false        // 插入时判断是否更新，更新时判断是否批量更新
  this.Distinct = ''      // 是否去重，返回一个[]string
  this.Condition = {} // update的时候查询条件
}

PostBody.prototype = {
  Get: function (table, data, orderby, limit, skip, distinct, select) {
    var body = new PostBody()
    body.Method = 'get'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    body.orderBy = orderby
    body.Limit = limit
    body.Skip = skip
    body.Distinct = distinct
    body.Select = select
    return body
  },
  Page: function (table, data, orderby, limit, skip, distinct, select) {
    var body = new PostBody()
    body.Method = 'page'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    body.orderBy = orderby
    body.Limit = limit
    body.Skip = skip
    body.Distinct = distinct
    body.Select = select
    return body
  },
  Count: function (table, data) {
    var body = new PostBody()
    body.Method = 'count'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    return body
  },
  Update: function (table, data, condition, updateAll) {
    var body = new PostBody()
    body.Method = 'update'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    body.UpdateAll = updateAll
    body.Condition = condition
    return body
  },
  Add: function (table, data, condition) {
    var body = new PostBody()
    body.Method = 'add'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    body.Condition = condition
    return body
  },
  Delete: function (table, data) {
    var body = new PostBody()
    body.Method = 'delete'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    return body
  }
}
