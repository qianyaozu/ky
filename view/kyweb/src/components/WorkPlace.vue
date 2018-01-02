<template>
  <div>
    <div class="tb-header">
      <span>工作面列表</span>
      <el-button  @click="addVisible = true" icon="el-icon-circle-plus-outline"  >新增</el-button>
      <el-button type="primary" @click="editOpen()" icon="el-icon-edit" >修改</el-button>
      <el-button type="danger" @click=" deleteOpen()" icon="el-icon-circle-close-outline">删除</el-button>
      <el-input placeholder="工作面名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>
      </el-input>
    </div>

    <el-table :data="DataList" @selection-change="SelectionChange" tooltipEffect="dark" :row-class-name="tableRowClassName">
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="ID"  label="编号"> </el-table-column>
      <el-table-column label="工作面名称" prop="Name">
      </el-table-column>
      <el-table-column prop="DipLength"  label="工作面斜长"> </el-table-column>
      <el-table-column prop="Capacity"  label="煤层容量"> </el-table-column>
      <el-table-column prop="Thickness"  label="煤层厚度"> </el-table-column>
      <el-table-column prop="Length"  label="回采总长"> </el-table-column>
      <el-table-column prop="create_at"  label="创建时间">
        <template slot-scope="scope">
          {{ format(scope.row.create_at) }}
        </template>
      </el-table-column>
    </el-table>
      <!--<el-pagination-->
        <!--layout="prev, pager, next"-->
        <!--:total="ItemsCount" :page-size="PageSize" v-on:current-change="Query()" :current-page.sync="CurrentPage">-->
      <!--</el-pagination>-->
    <el-dialog title="新增工作面" :visible.sync="addVisible"  width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-model="false" :close-on-click-modal="false">
            <el-form :model="AddForm"  class="demo-ruleForm"  label-width="100px">
              <el-row>
                <el-form-item label="工作面编号" prop="Name">
                  <el-input class="nameStyle" type="number" v-model="AddForm.ID" style="width: 85%;"></el-input>
                </el-form-item>
                <el-form-item label="工作面名称" prop="Name">
                  <el-input class="nameStyle" v-model="AddForm.Name" style="width: 85%;"></el-input>
                </el-form-item>
                <el-form-item label="工作面斜长" prop="Name">
                  <el-input class="nameStyle" type="number" v-model="AddForm.DipLength" style="width: 85%;"></el-input>
                  <el-label >m</el-label>
                </el-form-item>
                <el-form-item label="煤层容量" prop="Name">
                  <el-input class="nameStyle" type="number" v-model="AddForm.Capacity" style="width: 85%;"></el-input>
                  <el-label >m</el-label>
                </el-form-item>
                <el-form-item label="煤层厚度" prop="Name">
                  <el-input class="nameStyle" type="number" v-model="AddForm.Thickness" style="width: 85%;"></el-input>
                  <el-label >m</el-label>
                </el-form-item>
                <el-form-item label="回采总长" prop="Name">
                  <el-input class="nameStyle" type="number" v-model="AddForm.Length" style="width: 85%;"></el-input>
                  <el-label >m</el-label>
                </el-form-item>

              </el-row>
            </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addSubmit()">确 定</el-button>
        <el-button @click="addVisible = false">取 消</el-button>
      </span>
    </el-dialog>
    <el-dialog title="修改工作面" :visible.sync="editVisible"  width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-model="false" :close-on-click-modal="false">
      <el-form :model="EditForm"  class="demo-ruleForm"  label-width="100px">
        <el-row>
          <el-form-item label="工作面编号" prop="Name">
            <el-input class="nameStyle" type="number" v-model="EditForm.ID" :disabled="true" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="工作面名称" prop="Name">
            <el-input class="nameStyle" v-model="EditForm.Name" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="工作面斜长" prop="Name">
            <el-input class="nameStyle" type="number" v-model="EditForm.DipLength" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
          <el-form-item label="煤层容量" prop="Name">
            <el-input class="nameStyle" type="number" v-model="EditForm.Capacity" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
          <el-form-item label="煤层厚度" prop="Name">
            <el-input class="nameStyle" type="number" v-model="EditForm.Thickness" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
          <el-form-item label="回采总长" prop="Name">
            <el-input class="nameStyle" type="number" v-model="EditForm.Length" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>

        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="editSubmit()">确 定</el-button>
        <el-button @click="editVisible = false">取 消</el-button>
      </span>
    </el-dialog>

    <el-dialog title="删除" :visible.sync="deleteVisible" width="500px" center custom-class="danger-modal" :show-close="false">
      <span>确定要<span style="color: red">删除</span>选中数据么?</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="delSubmit()">确 定</el-button>
        <el-button @click="deleteVisible = false">取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
  import axios from 'axios'
  import PostBody from '../assets/js/PostBody'
  import {DateFormat} from '../assets/js/Common'
//  import UtilFormat from 'date-fns/format'
//  import {GetStore, SetStore} from '../vuex/store'
  export default {
    data () {
      return {
        TableName: 'WorkPlace',
        searchName: '',
        RowChecked: [],
        MultipleSelection: [],
        DataList: [],
        addVisible: false,
        deleteVisible: false,
        editVisible: false,
        AddForm: {
          ID: '',
          Name: '',
          DipLength: '',
          Capacity: '',
          Thickness: '',
          Length: ''
        },
        EditForm: {
          _id: '',
          ID: 0,
          Name: '',
          DipLength: 0,
          Capacity: 0,
          Thickness: 0,
          Length: ''
        },
        PageSize: 20
//        PageIndex: 0,
//        CurrentPage: 0,
//        ItemsCount: 0
      }
    },
    methods: {
      // format (ts) {
      //   var newDate = new Date()
      //   newDate.setTime(ts * 1000)
      //   var date = {
      //     'M+': newDate.getMonth() + 1,
      //     'd+': newDate.getDate(),
      //     'h+': newDate.getHours(),
      //     'm+': newDate.getMinutes(),
      //     's+': newDate.getSeconds(),
      //     'q+': Math.floor((newDate.getMonth() + 3) / 3),
      //     'S+': newDate.getMilliseconds()
      //   }
      //   var format = 'yyyy-MM-dd hh:mm:ss'
      //   if (/(y+)/i.test(format)) {
      //     format = format.replace(RegExp.$1, (newDate.getFullYear() + '').substr(4 - RegExp.$1.length))
      //   }
      //   for (var k in date) {
      //     if (new RegExp('(' + k + ')').test(format)) {
      //       format = format.replace(RegExp.$1, RegExp.$1.length === 1
      //         ? date[k] : ('00' + date[k]).substr(('' + date[k]).length))
      //     }
      //   }
      //   return format
      // },
      format (ts) {
        return DateFormat(ts)
      },
      validate (ob) {
        if (ob.ID === '') {
          this.$message.error('编号不能为空')
          return false
        }
        ob.ID = parseInt(ob.ID + '')
        if (ob.Name === '') {
          this.$message.error('工作地名称不能为空')
          return false
        }
        if (ob.DipLength === '') {
          this.$message.error('工作面斜长不能为空')
          return false
        }
        ob.DipLength = parseInt(ob.DipLength + '')
        if (ob.Capacity === '') {
          this.$message.error('煤矿容量不能为空')
          return false
        }
        ob.Capacity = parseInt(ob.Capacity + '')
        if (ob.Thickness === '') {
          this.$message.error('煤层厚度不能为空')
          return false
        }
        ob.Thickness = parseInt(ob.Thickness + '')
        if (ob.Length === '') {
          this.$message.error('回采总长不能为空')
          return false
        }
        ob.Length = parseInt(ob.Length + '')
        return true
      },
      tableRowClassName ({row, rowIndex}) {
        if (this.RowChecked.indexOf(rowIndex) >= 0) {
          return 'checked-bg'
        } else {
          return ''
        }
      },
      SelectionChange (selection) {
        this.RowChecked = []
        this.MultipleSelection = selection
        let ids = []
        this.MultipleSelection.map((s, j) => {
          ids.push(s.ID)
        })
        this.RowChecked = ids
      },
      delSubmit () {
        var query = {
          'ID': {
            '$in': this.RowChecked
          }
        }
        var body = new PostBody().Delete(this.TableName, query, '')
        axios.post('/api/common', body).then(response => {
          if (response.data.State === 0) {
            this.$message.error('删除工作面失败' + response.data.Message)
          } else {
            this.$message.success('删除成功')
            this.Query()
            this.deleteVisible = false
          }
        }).catch(err => {
          this.$message.error('删除工作面失败' + err)
        })
      },
      addSubmit () {
        if (this.AddForm.ID === 0) {
          this.$message.error('工作地编号不为空')
          return
        }
        if (!this.validate(this.AddForm)) {
          return
        }
        var condition = {
          '$or': [
            {
              'Name':
              this.AddForm.Name
            },
            {
              'ID':
              this.AddForm.ID
            }
          ]
        }
        var body = new PostBody().Add(this.TableName, this.AddForm, condition)
        axios.post('/api/common', body).then(response => {
          if (response.data.State === 0) {
            this.$message.error('新增工作面失败' + response.data.Message)
          } else {
            this.$message.success('新增成功')
            this.Query()
            this.addVisible = false
          }
        }).catch(err => {
          this.$message.error('新增工作面失败' + err)
        })
      },
      editOpen () {
        if (this.RowChecked.length === 0) {
          this.$message.error('请选中工作面')
          return
        }
        if (this.RowChecked.length > 1) {
          this.$message.error('请选中单个工作面')
          return
        }
        var row = this.MultipleSelection[0]
        this.EditForm._id = row._id
        this.EditForm.ID = row.ID
        this.EditForm.Name = row.Name
        this.EditForm.DipLength = row.DipLength
        this.EditForm.Capacity = row.Capacity
        this.EditForm.Thickness = row.Thickness
        this.EditForm.Length = row.Length
        this.editVisible = true
      },
      deleteOpen () {
        if (this.RowChecked.length === 0) {
          this.$message.error('请选中工作面')
          return
        }
        this.deleteVisible = true
      },
      editSubmit () {
        if (!this.validate(this.EditForm)) {
          return
        }
        var query = {
          'ID': this.EditForm.ID
        }

        var count = new PostBody().Count(this.TableName, {'Name': this.EditForm.Name, 'ID': {'$ne': this.EditForm.ID}})
        axios.post('/api/common', count).then(response => {
          if (response.data.State === 0) {
            this.$message.error('修改工作面失败' + response.data.Message)
          } else {
            if (response.data.Data === 0) {
              var body = new PostBody().Update(this.TableName, this.EditForm, query, false)
              axios.post('/api/common', body).then(response => {
                if (response.data.State === 0) {
                  this.$message.error('修改工作面失败' + response.data.Message)
                } else {
                  this.$message.success('修改成功')
                  this.Query()
                  this.editVisible = false
                }
              }).catch(err => {
                this.$message.error('修改工作面失败' + err)
              })
            } else {
              this.$message.error('该工作面名称已经被使用')
            }
          }
        }).catch(err => {
          this.$message.error('修改工作面失败' + err)
        })
      },
      searchSubmit () {
        this.Query()
      },
      Query () {
        var query = {}
        if (this.searchName !== '') {
          query = {
            'Name': {$regex: this.searchName, $options: 'i'}
          }
        }
        var body = new PostBody().Get(this.TableName, query, 'ID+', this.PageSize, 0, '', {})
        axios.post('/api/common', body).then(response => {
          this.DataList = response.data.Data
        })
      }
    },
    created () {
      this.Query()
    }
  }
</script>

