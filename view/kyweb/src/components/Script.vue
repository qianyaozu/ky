<template>
  <div>
    <div class="tb-header">
      <span>策略列表</span>
      <el-button class="btn-i-info" @click="addVisible = true" icon="el-icon-circle-plus-outline" style="color: #2C6D88;border:1px solid #2C6D88;font-size: 14px;margin-left: 10px">新增策略</el-button>
      <el-button class="btn-i-info" @click="editVisible = true" icon="el-icon-circle-plus-outline" style="color: #2C6D88;border:1px solid #2C6D88;font-size: 14px;margin-left: 10px">修改策略</el-button>
      <el-button class="btn-i-danger" @click=" deleteVisible = true" icon="el-icon-circle-close-outline">删除策略</el-button>
      <el-input placeholder="策略名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>
      </el-input>
    </div>

    <el-table :data="DataList" @selection-change="SelectionChange" tooltipEffect="dark" :row-class-name="tableRowClassName">
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="ID"  label="编号"> </el-table-column>
      <el-table-column label="策略名称" prop="Name">
        <template slot-scope="scope">
          <a href="#">{{ scope.row.Name }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="CreateTime"  label="创建时间"> </el-table-column>
    </el-table>

    <el-dialog title="新增策略" :visible.sync="addVisible"  width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-model="false" :close-on-click-modal="false">
            <el-form :model="AddForm"  class="demo-ruleForm"  label-width="100px">
              <el-row>
                <el-form-item label="策略名称" prop="Name">
                  <el-input class="nameStyle" v-model="AddForm.Name" style="width: 85%;"></el-input>
                </el-form-item>
                <el-form-item label="创建时间" prop="CreateTime">
                  <el-date-picker
                    v-model="AddForm.CreateTime"
                    type="date"
                    placeholder="选择日期"
                    value-format="yyyy-MM-dd" style="width: 85%">
                  </el-date-picker>
                </el-form-item>
              </el-row>
            </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addSubmit()">确 定</el-button>
        <el-button @click="addVisible = false">取 消</el-button>
      </span>
    </el-dialog>
    <el-dialog title="修改策略" :visible.sync="addVisible"  width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-model="false" :close-on-click-modal="false">
      <el-form :model="EditForm"  class="demo-ruleForm"  label-width="100px">
        <el-row>
          <el-form-item label="策略名称" prop="Name">
            <el-input class="nameStyle" v-model="EditForm.Name" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="创建时间" prop="CreateTime">
            <el-date-picker
              v-model="EditForm.CreateTime"
              type="date"
              placeholder="选择日期"
              value-format="yyyy-MM-dd" style="width: 85%">
            </el-date-picker>
          </el-form-item>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addSubmit()">确 定</el-button>
        <el-button @click="addVisible = false">取 消</el-button>
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
//  import {GetStore, SetStore} from '../vuex/store'
  export default {
    data () {
      return {
        searchName: '',
        RowChecked: [],
        MultipleSelection: [],
        DataList: [],
        addVisible: false,
        deleteVisible: false,
        editVisible: false,
        AddForm: {
          ID: 0,
          Name: '',
          CreateTime: '2018-01-01'
        },
        EditForm: {
          ID: 0,
          Name: '',
          CreateTime: '2018-01-01'
        }
      }
    },
    methods: {
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
        this.DataList.map((c, i) => {
          selection.map((s, j) => {
            if (c.ID === s.ID) {
              ids.push(i)
            }
          })
        })
        this.RowChecked = ids
      },
      delSubmit () {

      },
      addSubmit () {
        if (this.AddForm.Name === '') {
          this.$message.error('策略名称不为空')
        }
        axios.post('/api/addscript', this.AddForm).then(response => {
          if (response.data.State === 0) {
            this.$message.error('新增策略失败' + response.data.Message)
          } else {
            this.$message.success('成功')
          }
        }).catch(err => {
          this.$message.error('新增策略失败' + err)
        })
      },
      editSubmit () {
        if (this.EditForm.Name === '') {
          this.$message.error('策略名称不为空')
        }
        axios.post('/api/editscript', this.EditForm).then(response => {
          if (response.data.State === 0) {
            this.$message.error('修改策略失败' + response.data.Message)
          } else {
            this.$message.success('成功')
          }
        }).catch(err => {
          this.$message.error('修改策略失败' + err)
        })
      },
      searchSubmit () {
      }
    },
    created () {
      axios.get('/api/getScripts').then(response => {
        this.DataList = response.data.Data
      })
    }
  }
</script>
<style>
  .el-date-editor.el-input, .el-date-editor.el-input__inner{
    width: 85%;
  }
  .el-upload--text{
    width: 100%;
  }
  .el-button--small{
    height: 40px;
  }
</style>
