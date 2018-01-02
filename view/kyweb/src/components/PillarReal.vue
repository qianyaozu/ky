<template>
  <div>
    <div class="tb-header">
      <span>支柱列表</span>
      <el-input placeholder="支柱名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>

      </el-input>
      <el-select v-model="SearchFrameID" placeholder="请选择" style="float: right;margin-right: 20px;" >
        <el-option
          v-for="item in FrameList"
          :key="item.ID"
          :label="item.Name"
          :value="item.ID">
        </el-option>
      </el-select>
      <span style="float: right;margin-right: 20px;">支架</span>
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
      <el-table-column prop="CurrentPower"  label="当前阻力"> </el-table-column>
      <el-table-column prop="create_at"  label="创建时间">
        <template slot-scope="scope">
          {{ format(scope.row.create_at) }}
        </template>
      </el-table-column>
    </el-table>
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
        SearchFrameID: 0,
        TableName: 'PillarReal',
        searchName: '',
        RowChecked: [],
        MultipleSelection: [],
        DataList: [],
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
      searchSubmit() {
        this.Query()
      },
      Query() {
        var query = {}
        if (this.searchName !== '') {
          query.Name = {$regex: this.searchName, $options: 'i'}
        }
        if (this.SearchFrameID != 0) {
          query.FrameID = this.this.SearchFrameID
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
        var all = {ID: 0, Name: '全部'}
        this.FrameList.splice(0, 0, all)
        this.Query()
      })
    }
  }
</script>

