import Vue from 'vue'
import DataDict from '@/utils/dict'
import { listDictDataByCode as listDictDataByCode } from '@/api/system/dict/data'

function install() {
  Vue.use(DataDict, {
    metas: {
      '*': {
        labelField: 'label',
        valueField: 'value',
        request(dictMeta) {
          return listDictDataByCode(dictMeta.type).then(res => res.data.rows)
        },
      },
    },
  })
}

export default {
  install,
}