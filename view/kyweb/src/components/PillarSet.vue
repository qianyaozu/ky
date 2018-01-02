<!--ID              int //支柱编号-->
<!--PillarNo        int           //支柱编号-->
<!--PillarType      int           //1单柱 2多柱-->
<!--Name            string        //支柱名称-->
<!--FrameID         string        //支架编号-->
<!--MaxResistence   float32       //最大工作阻力-->
<!--InitPower       float32       //支柱初撑力-->
<!--DiameterDepth   float32       //缸径深度-->
<!--InstallPosition string-->
<template>
  <div>
    <div class="tb-header">
      <span>支柱列表</span>
      <el-button  @click="addVisible = true" icon="el-icon-circle-plus-outline"  >新增</el-button>
      <el-button type="primary" @click="editOpen()" icon="el-icon-edit" >修改</el-button>
      <el-button type="danger" @click=" deleteOpen()" icon="el-icon-circle-close-outline">删除</el-button>
      <el-input placeholder="支柱名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>
      </el-input>
    </div>

    <el-table :data="DataList" @selection-change="SelectionChange" tooltipEffect="dark" :row-class-name="tableRowClassName">
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="ID"  label="编号"> </el-table-column>
      <el-table-column label="支柱名称" prop="Name">
      </el-table-column>
      <el-table-column prop="InstallPosition"  label="安装位置"> </el-table-column>
      <el-table-column   label="支架">
        <template slot-scope="scope">
          {{ GetFrame(scope.row.FrameID) }}
        </template>
      </el-table-column>
      <el-table-column prop="MaxResistence"  label="最大工作阻力"> </el-table-column>
      <el-table-column prop="InitPower"  label="初撑力"> </el-table-column>
      <el-table-column prop="DiameterDepth"  label="缸径深度"> </el-table-column>
      <el-table-column prop="create_at"  label="创建时间">
        <template slot-scope="scope">
          {{ format(scope.row.create_at) }}
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="新增支柱" :visible.sync="addVisible"  width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-model="false" :close-on-click-modal="false">
      <el-form :model="AddForm"  class="demo-ruleForm"  label-width="100px">
        <el-row>
          <el-form-item label="支架编号" prop="ID">
            <el-input class="nameStyle" type="number" v-model="AddForm.ID" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="支架名称" prop="Name">
            <el-input class="nameStyle" v-model="AddForm.Name" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="安装位置" prop="InstallPosition">
            <el-input class="nameStyle" type="number" v-model="AddForm.InstallPosition" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
          <el-form-item label="所处z支架" prop="FrameID">
            <el-select v-model="AddForm.FrameID" placeholder="请选择">
              <el-option
                v-for="item in FrameList"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="最大工作阻力" prop="MaxResistence">
            <el-input class="nameStyle" type="number" v-model="AddForm.MaxResistence" style="width: 85%;"></el-input>
            <el-label >m3</el-label>
          </el-form-item>
          <el-form-item label="初撑力" prop="InitPower">
            <el-input class="nameStyle" type="number" v-model="AddForm.InitPower" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
            <el-form-item label="缸径深度" prop="DiameterDepth">
              <el-input class="nameStyle" type="number" v-model="AddForm.DiameterDepth" style="width: 85%;"></el-input>
              <el-label >m</el-label>
          </el-form-item>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addSubmit()">确 定</el-button>
        <el-button @click="addVisible = false">取 消</el-button>
      </span>
    </el-dialog>
    <el-dialog title="修改支柱" :visible.sync="editVisible"  width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-model="false" :close-on-click-modal="false">
      <el-form :model="EditForm"  class="demo-ruleForm"  label-width="100px">
        <el-row>
          <el-form-item label="支架编号" prop="Name">
            <el-input class="nameStyle" type="number" v-model="EditForm.ID" :disabled="true" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="支架名称" prop="Name">
            <el-input class="nameStyle" v-model="EditForm.Name" style="width: 85%;"></el-input>
          </el-form-item>
          <el-form-item label="安装位置" prop="InstallPosition">
            <el-input class="nameStyle" type="number" v-model="EditForm.InstallPosition" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
          <el-form-item label="所处工作面" prop="WorkPlaceID">
            <el-select v-model="EditForm.FrameID" placeholder="请选择">
              <el-option
                v-for="item in FrameList"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="最大工作阻力" prop="MaxResistence">
            <el-input class="nameStyle" type="number" v-model="EditForm.MaxResistence" style="width: 85%;"></el-input>
            <el-label >m3</el-label>
          </el-form-item>
          <el-form-item label="初撑力" prop="InitPower">
            <el-input class="nameStyle" type="number" v-model="EditForm.InitPower" style="width: 85%;"></el-input>
            <el-label >m</el-label>
          </el-form-item>
          <el-form-item label="缸径深度" prop="DiameterDepth">
            <el-input class="nameStyle" type="number" v-model="EditForm.DiameterDepth" style="width: 85%;"></el-input>
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
  //  import {GetStore, SetStore} from '../vuex/store'
  export default {
    data() {
      return {
        TableName: 'PillarSet',
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
          InstallPosition: '',
          InitPower: 0,
          MaxResistence: 0,
          DiameterDepth: 0,
          FrameID: null
        },
        EditForm: {
          ID: 0,
          Name: '',
          InstallPosition: 0,
          InitPower: 0,
          MaxResistence: 0,
          DiameterDepth: 0,
          FrameID: null
        },
        FrameList: []
      }
    },
    methods: {
      format(ts) {
        return DateFormat(ts)
      },
      GetFrame(id) {
        for (var i = 0; i < this.FrameList.length; i++) {
          if (this.FrameList[i].ID === id) {
            return this.FrameList[i].Name
          }
        }
      },
      validate(ob) {
        if (ob.ID === '') {
          this.$message.error('编号不能为空')
          return false
        }
        ob.ID = parseInt(ob.ID + '')
        if (ob.Name === '') {
          this.$message.error('支柱名称不能为空')
          return false
        }
        if (ob.InitPower === '') {
          this.$message.error('初撑力不能为空')
          return false
        }
        ob.InitPower = parseInt(ob.InitPower + '')
        if (ob.MaxResistence === '') {
          this.$message.error('最大工作阻力不能为空')
          return false
        }
        ob.MaxResistence = parseInt(ob.MaxResistence + '')
        return true
      },
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
        var body = new PostBody().Delete(this.TableName, query, '')
        axios.post('/api/common', body).then(response => {
          if (response.data.State === 0) {
            this.$message.error('删除支柱失败' + response.data.Message)
          } else {
            this.$message.success('删除成功')
            this.Query()
            this.deleteVisible = false
          }
        }).catch(err => {
          this.$message.error('删除支柱失败' + err)
        })
      },
      addSubmit() {
        if (this.AddForm.ID === 0) {
          this.$message.error('支柱编号不为空')
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
            this.$message.error('新增支柱失败' + response.data.Message)
          } else {
            this.$message.success('新增成功')
            this.Query()
            this.addVisible = false
          }
        }).catch(err => {
          this.$message.error('新增支柱失败' + err)
        })
      },
      editOpen() {
        if (this.RowChecked.length === 0) {
          this.$message.error('请选中支柱')
          return
        }
        if (this.RowChecked.length > 1) {
          this.$message.error('请选中单个支柱')
          return
        }
        var row = this.MultipleSelection[0]
        this.EditForm._id = row._id
        this.EditForm.ID = row.ID
        this.EditForm.Name = row.Name
        this.EditForm.InstallPosition = row.InstallPosition
        this.EditForm.InitPower = row.InitPower
        this.EditForm.MaxResistence = row.MaxResistence
        this.EditForm.FrameID = row.FrameID
        this.EditForm.DiameterDepth = row.DiameterDepth
        this.editVisible = true
      },
      deleteOpen() {
        if (this.RowChecked.length === 0) {
          this.$message.error('请选中支柱')
          return
        }
        this.deleteVisible = true
      },
      editSubmit() {
        if (!this.validate(this.EditForm)) {
          return
        }
        var query = {
          'ID': this.EditForm.ID
        }

        var count = new PostBody().Count(this.TableName, {
          'Name': this.EditForm.Name,
          'ID': {'$ne': this.EditForm.ID}
        })
        axios.post('/api/common', count).then(response => {
          if (response.data.State === 0) {
            this.$message.error('修改失败' + response.data.Message)
          } else {
            if (response.data.Data === 0) {
              var body = new PostBody().Update(this.TableName, this.EditForm, query, false)
              axios.post('/api/common', body).then(response => {
                if (response.data.State === 0) {
                  this.$message.error('修改失败' + response.data.Message)
                } else {
                  this.$message.success('修改成功')
                  this.Query()
                  this.editVisible = false
                }
              }).catch(err => {
                this.$message.error('修改失败' + err)
              })
            } else {
              this.$message.error('该名称已经被使用')
            }
          }
        }).catch(err => {
          this.$message.error('修改失败' + err)
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
        var body = new PostBody().Get(this.TableName, query, 'ID+', 0, 0, '', {})
        axios.post('/api/common', body).then(response => {
          this.DataList = response.data.Data
        })
      }
    },
    created() {
      var body = new PostBody().Get('FrameSet', {}, '', 0, 0, '', {})
      axios.post('/api/common', body).then(response => {
        this.FrameList = response.data.Data
        this.Query()
      })
    }
  }
</script>

