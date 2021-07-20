<template>
  <div class="container">
  <form class="form-horizontal" role="form">
    <div class="form-group">
      <label for="name" class="col-sm-2 control-label">名字</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" id="name" v-model="user.name"
             placeholder="请输入名字">
      </div>
    </div>
    <div class="form-group">
      <label for="sex" class="col-sm-2 control-label">性别</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" id="sex" v-model="user.sex"
             placeholder="请输入名字">
      </div>
    </div>
    <div class="form-group">
      <label for="age" class="col-sm-2 control-label">年龄</label>
      <div class="col-sm-10">
        <input type="text" class="form-control" id="age" v-model="user.age"
             placeholder="请输入名字">
      </div>
    </div>
    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <a v-on:click="addUser">添加</a>
      </div>
    </div>
  </form>
  <table class="table table-bordered">
    <caption>人员列表</caption>
    <thead>
      <tr>        
        <th>序号</th>
        <th>姓名</th>
        <th>性别</th>
        <th>年龄</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(user,index) in users" :key="user.name">
        <td>{{(index+1)}}</td>
        <td>{{user.name}}</td>
        <td>{{user.sex}}</td>
        <td>{{user.age}}</td>
        <td><a v-on:click="delUser(index)">删除</a></td>
      </tr>
      <tr><td colspan="5"><a v-on:click="delUser(-1)">删除全部</a></td></tr>
    </tbody>
  </table>
</div>
</template>

<script>
export default {
  el: 'container',
  data() {
    return {
    user: {
      name: null,
      sex: null,
      age: null,
    },
    users: [{name: "csf", sex: "男", age: 20}]}
  },
  methods: {
    addUser(){
      this.users.push(this.user);
      this.user={};
    },
    delUser(us_index){
      if(us_index == -1){
          this.users={};
      }else{
          this.users.splice(us_index,1);
      }        
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
</style>
