<template>
  <div>

    <router-view />
</div>
</template>

<script>
  import {GuestHttp} from '@/axios/api.js'
import Cookies from 'js-cookie'

export default {
  data() {
    return {
     
    }
  },

  created(){
    
    // this.init()
  },
  methods:{
   
  }
};
</script>