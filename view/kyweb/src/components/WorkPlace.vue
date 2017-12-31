<template>
  <div>
    <div class="tb-header">
      <span>工作面列表</span>
      <el-button class="btn-i-info" @click="addVisible = true" icon="el-icon-circle-plus-outline" style="color: #2C6D88;border:1px solid #2C6D88;font-size: 14px;margin-left: 10px">新增</el-button>
      <el-button class="btn-i-info" @click="editOpen()" icon="el-icon-circle-plus-outline" style="color: #2C6D88;border:1px solid #2C6D88;font-size: 14px;margin-left: 10px">修改</el-button>
      <el-button class="btn-i-danger" @click=" deleteOpen()" icon="el-icon-circle-close-outline">删除</el-button>
      <el-input placeholder="工作面名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>
      </el-input>
    </div>

    <el-table :data="DataList" @selection-change="SelectionChange" tooltipEffect="dark" :row-class-name="tableRowClassName">
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="ID"  label="编号"> </el-table-column>
      <el-table-column label="工作面名称" prop="Name">
        <template slot-scope="scope">
          <a href="#">{{ scope.row.Name }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="DipLength"  label="工作面斜长"> </el-table-column>
      <el-table-column prop="Capacity"  label="煤层容量"> </el-table-column>
      <el-table-column prop="Thickness"  label="煤层厚度"> </el-table-column>
      <el-table-column prop="Length"  label="回采总长"> </el-table-column>
      <el-table-column prop="create_at"  label="创建时间"> </el-table-column>
    </el-table>

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
                  <el-label >m3</el-label>
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
            <el-label >m3</el-label>
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
//  import {GetStore, SetStore} from '../vuex/store'
  export default {
    data() {
      return {
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
        }
      }
    },
    methods: {
      tableRowClassName({row, rowIndex}) {
        if (this.RowChecked.indexOf(rowIndex) >= 0) {
          return 'checked-bg'
        } else {
          return ''
        }
      },
      SelectionChange(selection) {
        this.RowChecked = []
        this.MultipleSelection = selection
        let ids = []
        this.MultipleSelection.map((s, j) => {
          ids.push(s.ID)
        })
        this.RowChecked = ids
      },
      delSubmit() {
        var query = {
          'ID': {
            '$in': this.RowChecked
          }
        }
        var body = new PostBody().Delete('workplace', query, '')
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
      addSubmit() {
        if (this.AddForm.ID === 0) {
          this.$message.error('工作地编号不为空')
          return
        }
        if (this.AddForm.Name === '') {
          this.$message.error('工作地名称不为空')
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
        var body = new PostBody().Add('workplace', this.AddForm, condition)
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
      editOpen() {
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
      deleteOpen() {
        if (this.RowChecked.length === 0) {
          this.$message.error('请选中工作面')
          return
        }
        this.deleteVisible = true
      },
      editSubmit() {
        if (this.EditForm.Name === '') {
          this.$message.error('策略名称不为空')
        }
        var query = {
          'ID': this.EditForm.ID
        }

        var count = new PostBody().Count('workplace', {'Name': this.EditForm.Name, 'ID': {'$ne': this.EditForm.ID}})
        axios.post('/api/common', count).then(response => {
          if (response.data.State === 0) {
            this.$message.error('修改工作面失败' + response.data.Message)
          } else {
            if(response.data.Data === 0) {
              var body = new PostBody().Update('workplace', this.EditForm, query, false)
              axios.post('/api/common', body).then(response => {
                if (response.data.State === 0) {
                  this.$message.error('修改工作面失败' + response.data.Message)
                } else {
                  this.$message.success('修改成功')
                  this.Query()
                  this.editVisible = false
                }
              }).catch(err => {
                this.$message.error('修改策略失败' + err)
              })
            }else{
              this.$message.error('该策略名称已经被使用')
            }
          }
        }).catch(err => {
          this.$message.error('修改策略失败' + err)
        })
      },
      searchSubmit() {
        this.Query()
      },
      Query() {
        var query = {}
        if (this.searchName !== '') {
          query = {
            'Name': {$regex: this.searchName, $options: 'i'}
          }
        }
        var body = new PostBody().Get('workplace', query, 'ID+', 20, 0, '', {})
        axios.post('/api/common', body).then(response => {
          this.DataList = response.data.Data
        })
      }
    },
    created() {
      this.Query()
    }
  }
</script>

