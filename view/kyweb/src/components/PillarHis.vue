<template>
  <div>
    <div class="tb-header">
      <span>历史支柱</span>
      <el-input placeholder="支柱名称"  class="input-with-select" style="" v-model="searchName" @keyup.enter.native="searchSubmit">
        <!--<i class="el-icon-search" slot="append"></i>-->
        <el-button slot="append" icon="el-icon-search" style="color: rgba(25, 158, 216, 1);" @click="searchSubmit"></el-button>

      </el-input>
      <el-select v-model="SearchFrameID" placeholder="请选择" class="floatright" v-on:change="FrameChanged" >
        <el-option
          v-for="item in FrameList"
          :key="item.ID"
          :label="item.Name"
          :value="item.ID">
        </el-option>
      </el-select>
      <span class="floatright">支架</span>

      <el-date-picker class="floatright"
        v-model="SearchDate" @change="SearchDateChanged"
        type="daterange"
        align="right"
        unlink-panels
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        :picker-options="pickerOptions2">
      </el-date-picker>
      <span class="floatright">查询时间</span>
    </div>
    <el-table :data="DataList" @selection-change="SelectionChange" tooltipEffect="dark" :row-class-name="tableRowClassName">
      <el-table-column prop="ID"  label="编号"> </el-table-column>
      <el-table-column label="支柱名称" prop="Name">
      </el-table-column>
      <el-table-column prop="InstallPosition"  label="安装位置"> </el-table-column>
      <el-table-column   label="支架">
        <template slot-scope="scope">
          {{ GetFrame(scope.row.FrameID) }}
        </template>
      </el-table-column>
      <el-table-column prop="CurrentPower"  label="当前阻力"> </el-table-column>
      <el-table-column prop="create_at"  label="创建时间">
        <template slot-scope="scope">
          {{ format(scope.row.create_at) }}
        </template>
      </el-table-column>
    </el-table>
    <el-pagination class="floatright"
                   @current-change="handleCurrentChange"
                   :current-page.sync="CurrentPage"
                   :page-size="PageSize"
                   layout="total, prev, pager, next"
                   :total="DataTotal">
    </el-pagination>
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
        DataTotal: 0,
        PageSize: 100,
        CurrentPage: 1,
        SearchDate: [],
        SearchFrameID: 0,
        TableName: 'PillarReal',
        searchName: '',
        RowChecked: [],
        MultipleSelection: [],
        DataList: [],
        FrameList: [],
        pickerOptions2: {
          shortcuts: [{
            text: '最近一周',
            onClick(picker) {
              const end = new Date()
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
              picker.$emit('pick', [start, end])
            }
          }, {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date()
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
              picker.$emit('pick', [start, end])
            }
          }, {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date()
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
              picker.$emit('pick', [start, end])
            }
          }]
        },
      }
    },
    methods: {
      handleCurrentChange() {
        this.Query()
      },
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
        this.CurrentPage = 1
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
        if (this.SearchDate != '') {
          query.create_at = {'$gt': this.SearchDate[0]}
          query.create_at = {'$lt': this.SearchDate[1]}
        }
        var skip = 0
        if (this.CurrentPage > 1) {
          skip = (this.CurrentPage - 1) * this.PageSize
        }
        var body = new PostBody().Page(this.TableName, query, 'ID+', this.PageSize, skip, '', {})
        axios.post('/api/common', body).then(response => {
          this.DataList = response.data.Data.Items
          this.DataTotal = response.data.Data.Count
        })
      },
      SearchDateChanged() {
        this.CurrentPage = 1
        alert(this.SearchDate)
        this.Query()
      },
      FrameChanged() {
        this.CurrentPage = 1
        this.Query()
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

