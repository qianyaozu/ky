<template>
  <div>
    <div class="tb-header">
      <span>实时支架</span>
     <el-input placeholder="支架名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>
      </el-input>
    </div>

    <el-table :data="DataList" @selection-change="SelectionChange" tooltipEffect="dark" :row-class-name="tableRowClassName">
      <el-table-column type="selection"></el-table-column>
      <el-table-column prop="ID"  label="编号"> </el-table-column>
      <el-table-column label="支架名称" prop="Name">
      </el-table-column>
      <el-table-column prop="InstallPosition"  label="安装位置"> </el-table-column>
      <el-table-column   label="工作面">
        <template slot-scope="scope">
          {{ GetWorkPlace(scope.row.WorkPlaceID) }}
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
    data () {
      return {
        TableName: 'FrameReal',
        searchName: '',
        RowChecked: [],
        MultipleSelection: [],
        DataList: [],
        addVisible: false,
        deleteVisible: false,
        editVisible: false,
        WorkPlaceList: []
      }
    },
    methods: {
      format (ts) {
        return DateFormat(ts)
      },
      GetWorkPlace (id) {
        for (var i = 0; i < this.WorkPlaceList.length; i++) {
          if (this.WorkPlaceList[i].ID === id) {
            return this.WorkPlaceList[i].Name
          }
        }
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
        var body = new PostBody().Get(this.TableName, query, 'ID+', 0, 0, '', {})
        axios.post('/api/common', body).then(response => {
          this.DataList = response.data.Data
        })
      }
    },
    created () {
      var body = new PostBody().Get('WorkPlace', {}, '', 0, 0, '', {})
      axios.post('/api/common', body).then(response => {
        this.WorkPlaceList = response.data.Data
        this.Query()
      })
    }
  }
</script>

