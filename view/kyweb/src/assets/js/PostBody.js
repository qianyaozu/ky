function PostBody () {
  this.Method = ''  // get,count,update,exists,add,delete,func
  this.DBName = ''
  this.Table = ''
  this.Data = ''
  this.OrderBy = ''
  this.Limit = 0
  this.Skip = 0
  this.Select = ''
  this.ObjectID = ''
  this.UpdateAll = false        // 插入时判断是否更新，更新时判断是否批量更新
  this.Distinct = ''      // 是否去重，返回一个[]string
  this.Condition = '' // update的时候查询条件
}

PostBody.prototype = {
  Get: function (table, data, orderby, limit, skip, obid, distinct) {
    var body = new PostBody()
    body.Method = 'get'
    body.DBName = 'ky'
    body.Table = table
    body.Data = data
    body.orderBy = orderby
    body.Limit = limit
    body.Skip = skip
    body.ObjectID = obid
    body.Distinct = distinct
    return body
  }
}
