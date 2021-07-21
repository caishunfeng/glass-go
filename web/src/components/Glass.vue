<template>
  <div class="container">
  <table class="pure-table pure-table-bordered">
    <caption>域名列表</caption>
    <thead>
      <tr>        
        <th>序号</th>
        <th>代理域名</th>
        <th>目标域名</th>
        <th>过期时间</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(domain,index) in domains" :key="index">
        <td>{{(index+1)}}</td>
        <td v-if="!domain.isNew">{{domain.from}}</td>
        <td v-else><input v-model="domain.from" placeholder="代理域名"></td>
        <td v-if="!domain.isNew">{{domain.to}}</td>
        <td v-else><input v-model="domain.to" placeholder="目标域名"></td>
        <td v-if="!domain.isNew" type="number">{{domain.expire}}</td>
        <td v-else><input type="number" v-model="domain.expire" placeholder="过期时间"></td>
        <td v-if="!domain.isNew" v-on:click="cancel(index)">删除</td>
        <td v-else v-on:click="registry(domain)">注册</td>
        <!-- <div v-else>
          <td>{{(index+1)}}</td>
          <input v-model="domain.from" placeholder="请输入代理域名">
          <input v-model="domain.to" placeholder="请输入目标域名">
          <input type="number" v-model="domain.expire" placeholder="请输入过期时间">
          <a v-on:click="registry(index)">注册</a>
        </div> -->
      </tr>
      <tr>
        <td colspan="5"><button v-on:click="newDomain">新增</button></td>
      </tr>
    </tbody>
  </table>
</div>
</template>

<script>
import axios from 'axios'
export default {
  el: 'container',
  data() {
    return {
      domain: {
        from: null,
        to: null,
        expire: null,
        isNew: true,
      },
      domains: []
    }
  },
  created() {
    console.log("ready go");
      this.getAllProxyDomains();
  },
  methods: {
    getAllProxyDomains() {
      axios.get('http://localhost:8888/getProxyKeys')
      .then(response => {
        console.log('getProxyKeys resp:', response);
        this.domains = response.data.data;
      })
      .catch(response => {
        console.error(response)
      })
    },
    registry(domain){
      axios.get(`http://localhost:8888/registry?from=${domain.from}&to=${domain.to}&expire=${domain.expire}`)
      .then(response => {
        console.log('registry resp:', response);
        this.getAllProxyDomains();
      })
      .catch(response => {
        console.error(response)
      });
    },
    cancel(index){
      axios.get(`http://localhost:8888/cancel?from=${this.domains[index].from}`)
      .then(response => {
        console.log('cancel resp:', response);
        this.getAllProxyDomains();
      })
      .catch(response => {
        console.error(response)
      });
    },
    newDomain(){
      this.domains.push({
        from: null,
        to: null,
        expire: null,
        isNew: true,
      });
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
html {
    font-family: sans-serif;
    -ms-text-size-adjust: 100%;
    -webkit-text-size-adjust: 100%;
}
 
body {
    margin: 10px;
}
table {
    border-collapse: collapse;
    border-spacing: 0;
}
 
td,th {
    padding: 0;
}
 
.pure-table {
    border-collapse: collapse;
    border-spacing: 0;
    empty-cells: show;
    border: 1px solid #cbcbcb;
}
 
.pure-table caption {
    color: #000;
    font: italic 85%/1 arial,sans-serif;
    padding: 1em 0;
    text-align: center;
}
 
.pure-table td,.pure-table th {
    border-left: 1px solid #cbcbcb;
    border-width: 0 0 0 1px;
    font-size: inherit;
    margin: 0;
    overflow: visible;
    padding: .5em 1em;
}
 
.pure-table thead {
    background-color: #e0e0e0;
    color: #000;
    text-align: left;
    vertical-align: bottom;
}
 
.pure-table td {
    background-color: transparent;
}
 
.pure-table-bordered td {
    border-bottom: 1px solid #cbcbcb;
}
 
.pure-table-bordered tbody>tr:last-child>td {
    border-bottom-width: 0;
}
</style>
