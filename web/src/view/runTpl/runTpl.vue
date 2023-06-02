<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="模板类型:">
          <el-select
            v-model="searchInfo.tplType"
            placeholder="请选择模板类型"
            filterable
            style="width: 120px"
            clearable
          >
            <el-option v-for="(item,key) in tplTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-form-item label="模板类型:">
          <el-select
            v-model="formData.tplType"
            placeholder="请选择模板类型"
            filterable
            style="width: 160px"
            clearable
          >
            <el-option v-for="(item,key) in tplTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-button type="primary" style="margin-left: 10px;" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="110">
          <template #default="scope">{{ formatDateShort(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="页面ID" prop="tplId" width="100" /> -->
        <el-table-column align="left" label="页面名" prop="tplName" width="200" />
        <el-table-column align="left" label="模板类型" prop="tplType" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.tplType, tplTypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="点击链接" prop="clickUrl" width="350" />
        <el-table-column align="left" label="展示链接" prop="pageUrl" width="350" />
        <!-- <el-table-column align="left" label="按钮描述" prop="clickDesc" width="120" />
        <el-table-column align="left" label="弹窗描述" prop="dialogDesc" width="120" />
        <el-table-column align="left" label="描述1" prop="text1" width="120" />
        <el-table-column align="left" label="描述2" prop="text2" width="120" />
        <el-table-column align="left" label="描述3" prop="text3" width="120" />
        <el-table-column align="left" label="描述4" prop="text4" width="120" />
        <el-table-column align="left" label="描述5" prop="text5" width="120" />
        <el-table-column align="left" label="描述6" prop="text6" width="120" />
        <el-table-column align="left" label="描述7" prop="text7" width="120" />
        <el-table-column align="left" label="描述8" prop="text8" width="120" />
        <el-table-column align="left" label="描述9" prop="text9" width="120" />
        <el-table-column align="left" label="描述10" prop="text10" width="120" />
        <el-table-column align="left" label="描述11" prop="text11" width="120" />
        <el-table-column align="left" label="描述12" prop="text12" width="120" />
        <el-table-column align="left" label="描述13" prop="text13" width="120" />
        <el-table-column align="left" label="描述14" prop="text14" width="120" />
        <el-table-column align="left" label="描述15" prop="text15" width="120" />
        <el-table-column align="left" label="描述16" prop="text16" width="120" />
        <el-table-column align="left" label="描述17" prop="text17" width="120" />
        <el-table-column align="left" label="描述18" prop="text18" width="120" />
        <el-table-column align="left" label="描述19" prop="text19" width="120" />
        <el-table-column align="left" label="描述20" prop="text20" width="120" />
        <el-table-column align="left" label="图片1" prop="picName1" width="120" />
        <el-table-column align="left" label="图片2" prop="picName2" width="120" />
        <el-table-column align="left" label="图片3" prop="picName3" width="120" />
        <el-table-column align="left" label="图片4" prop="picName4" width="120" />
        <el-table-column align="left" label="图片5" prop="picName5" width="120" />
        <el-table-column align="left" label="图片6" prop="picName6" width="120" />
        <el-table-column align="left" label="图片7" prop="picName7" width="120" />
        <el-table-column align="left" label="图片8" prop="picName8" width="120" />
        <el-table-column align="left" label="图片9" prop="picName9" width="120" />
        <el-table-column align="left" label="图片10" prop="picName10" width="120" /> -->
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateRunTplFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="Plus" @click="copyRow(scope.row)">复制</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="110px">
        <!-- <el-form-item label="模板页面ID:" prop="tplId">
          <el-input v-model="formData.tplId" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="模板网页名:" prop="tplName">
          <el-input v-model="formData.tplName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="点击链接:" prop="clickUrl">
          <el-input v-model="formData.clickUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <div v-show="!isOpenStock">
          <el-form-item label="按钮描述:" prop="clickDesc">
            <el-input v-model="formData.clickDesc" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="弹窗描述:" prop="dialogDesc">
            <el-input v-model="formData.dialogDesc" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <!-- <el-form-item label="展示链接:" prop="pageUrl">
            <el-input v-model="formData.pageUrl" :clearable="true" placeholder="请输入" />
          </el-form-item> -->
          <el-form-item label="描述1:" prop="text1">
            <el-input v-model="formData.text1" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述2:" prop="text2">
            <el-input v-model="formData.text2" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述3:" prop="text3">
            <el-input v-model="formData.text3" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述4:" prop="text4">
            <el-input v-model="formData.text4" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述5:" prop="text5">
            <el-input v-model="formData.text5" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述6:" prop="text6">
            <el-input v-model="formData.text6" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述7:" prop="text7">
            <el-input v-model="formData.text7" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述8:" prop="text8">
            <el-input v-model="formData.text8" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述9:" prop="text9">
            <el-input v-model="formData.text9" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述10:" prop="text10">
            <el-input v-model="formData.text10" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述11:" prop="text11">
            <el-input v-model="formData.text11" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述12:" prop="text12">
            <el-input v-model="formData.text12" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述13:" prop="text13">
            <el-input v-model="formData.text13" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述14:" prop="text14">
            <el-input v-model="formData.text14" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述15:" prop="text15">
            <el-input v-model="formData.text15" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <!-- <el-form-item label="描述16:" prop="text16">
            <el-input v-model="formData.text16" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述17:" prop="text17">
            <el-input v-model="formData.text17" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述18:" prop="text18">
            <el-input v-model="formData.text18" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述19:" prop="text19">
            <el-input v-model="formData.text19" :clearable="true" placeholder="请输入" />
          </el-form-item>
          <el-form-item label="描述20:" prop="text20">
            <el-input v-model="formData.text20" :clearable="true" placeholder="请输入" />
          </el-form-item> -->

          <el-form-item label="图片1-1080x675:" prop="picName1" label-position="left" label-width="150px">
            <upload-common @on-success="upImg1" />
            <el-button v-show="formData.picName1 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName1)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName1 !== ''" pic-type="file" :pic-src="formData.picName1" />
          <br>
          <el-form-item label="图片2-350x400:" prop="picName2" label-position="left" label-width="150px">
            <upload-common @on-success="upImg2" />
            <el-button v-show="formData.picName2 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName2)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName2 !== ''" pic-type="file" :pic-src="formData.picName2" />
          <br>
          <el-form-item label="图片3-350x400:" prop="picName1" label-position="left" label-width="150px">
            <upload-common @on-success="upImg3" />
            <el-button v-show="formData.picName3 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName3)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName3 !== ''" pic-type="file" :pic-src="formData.picName3" />
          <br>
          <el-form-item label="图片4-350x400:" prop="picName4" label-position="left" label-width="150px">
            <upload-common @on-success="upImg4" />
            <el-button v-show="formData.picName4 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName4)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName4 !== ''" pic-type="file" :pic-src="formData.picName4" />
          <br>
          <el-form-item label="图片5-350x400:" prop="picName5" label-position="left" label-width="150px">
            <upload-common @on-success="upImg5" />
            <el-button v-show="formData.picName5 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName5)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName5 !== ''" pic-type="file" :pic-src="formData.picName5" />
          <br>
          <el-form-item label="图片6-350x400:" prop="picName1" label-position="left" label-width="150px">
            <upload-common @on-success="upImg6" />
            <el-button v-show="formData.picName6 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName6)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName6 !== ''" pic-type="file" :pic-src="formData.picName6" />
          <br>
          <el-form-item label="图片7-350x400:" prop="picName7" label-position="left" label-width="150px">
            <upload-common @on-success="upImg7" />
            <el-button v-show="formData.picName7 !== ''" style="margin-left: 20px;" type="primary" link icon="delete" @click="deleteImg(formData.picName7)">删除</el-button>
          </el-form-item>
          <CustomPic v-show="formData.picName7 !== ''" pic-type="file" :pic-src="formData.picName7" />
          <br>
        </div>
        <!-- <el-form-item label="图片7:" prop="picName7">
          <el-input v-model="formData.picName7" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片8:" prop="picName8">
          <el-input v-model="formData.picName8" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片9:" prop="picName9">
          <el-input v-model="formData.picName9" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片10:" prop="picName10">
          <el-input v-model="formData.picName10" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'RunTpl'
}
</script>

<script setup>
import {
  createRunTpl,
  deleteRunTpl,
  deleteRunTplByIds,
  copyRunTpl,
  updateRunTpl,
  findRunTpl,
  getRunTplList
} from '@/api/runTpl'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatDateShort, formatBoolean, filterDict } from '@/utils/format'
import { deleteFileByUrl } from '@/api/fileUploadAndDownload'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import UploadCommon from '@/components/upload/common.vue'
import CustomPic from '@/components/customPic/index.vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  tplId: '',
  tplName: '',
  clickUrl: '',
  clickDesc: '',
  dialogDesc: '',
  pageUrl: '',
  text1: '',
  text2: '',
  text3: '',
  text4: '',
  text5: '',
  text6: '',
  text7: '',
  text8: '',
  text9: '',
  text10: '',
  text11: '',
  text12: '',
  text13: '',
  text14: '',
  text15: '',
  text16: '',
  text17: '',
  text18: '',
  text19: '',
  text20: '',
  picName1: '',
  picName2: '',
  picName3: '',
  picName4: '',
  picName5: '',
  picName6: '',
  picName7: '',
  picName8: '',
  picName9: '',
  picName10: '',
})

// 验证规则
const rule = reactive({
  tplName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const tplTypeOptions = ref([])
const isOpenStock = ref(false)
const path = ref(import.meta.env.VITE_BASE_API + '/')

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getRunTplList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteImg = async(name) => {
  console.log('delete name --- ' + name)
  const res = await deleteFileByUrl({ url: name })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功!',
    })
    Object.keys(formData.value).forEach(key => {
    // 如果值为 "nvgoreigjeorigj"，则将其设置为空
      if (formData.value[key] === name) {
        formData.value[key] = ''
      }
    })
  }
}
const upImg1 = async(url) => {
  formData.value.picName1 = url
}
const upImg2 = async(url) => {
  formData.value.picName2 = url
}
const upImg3 = async(url) => {
  formData.value.picName3 = url
}
const upImg4 = async(url) => {
  formData.value.picName4 = url
}
const upImg5 = async(url) => {
  formData.value.picName5 = url
}
const upImg6 = async(url) => {
  formData.value.picName6 = url
}
const upImg7 = async(url) => {
  formData.value.picName7 = url
}
// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  tplTypeOptions.value = await getDictFunc('tplType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteRunTplFunc(row)
  })
}

const copyRow = (row) => {
  ElMessageBox.confirm('确定要复制吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    copyRunTplFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteRunTplByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateRunTplFunc = async(row) => {
  const res = await findRunTpl({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rerunTpl
    dialogFormVisible.value = true
    if (formData.value.tplType === 2 || formData.value.tplType === 3) {
      isOpenStock.value = true
    }
  }
}

// 删除行
const deleteRunTplFunc = async(row) => {
  const res = await deleteRunTpl({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 复制行
const copyRunTplFunc = async(row) => {
  const res = await copyRunTpl({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '复制成功'
    })
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  if (!formData.value.tplType) {
    ElMessage({
      type: 'error',
      message: '请先选择模板类型'
    })
    return
  }
  isOpenStock.value = false
  if (formData.value.tplType === 1) {
    type.value = 'create'
    dialogFormVisible.value = true
    console.log('创建模板1-----------')
  } else if (formData.value.tplType === 2 || formData.value.tplType === 3) {
    type.value = 'create'
    dialogFormVisible.value = true
    isOpenStock.value = true
  } else {
    ElMessage({
      type: 'info',
      message: '不支持'
    })
  }
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    tplId: '',
    tplName: '',
    clickUrl: '',
    clickDesc: '',
    dialogDesc: '',
    pageUril: '',
    text1: '',
    text2: '',
    text3: '',
    text4: '',
    text5: '',
    text6: '',
    text7: '',
    text8: '',
    text9: '',
    text10: '',
    text11: '',
    text12: '',
    text13: '',
    text14: '',
    text15: '',
    text16: '',
    text17: '',
    text18: '',
    text19: '',
    text20: '',
    picName1: '',
    picName2: '',
    picName3: '',
    picName4: '',
    picName5: '',
    picName6: '',
    picName7: '',
    picName8: '',
    picName9: '',
    picName10: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createRunTpl(formData.value)
           break
         case 'update':
           res = await updateRunTpl(formData.value)
           break
         default:
           res = await createRunTpl(formData.value)
           break
       }
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: '创建/更改成功'
         })
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>
