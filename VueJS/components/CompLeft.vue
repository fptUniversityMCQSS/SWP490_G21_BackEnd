<template>
  <div class="col-md-2 p-x-0 p-y-3 leftmenu">
    <ul class="admin-sidenav list-unstyled">
      <li>
        <router-link to="/q&a">Question and Answer</router-link>
      </li>
      <li>
        <router-link to="/listknow">List of knowledges</router-link>
      </li>
      <li>
        <router-link to="/uploadknow">Upload knowledges</router-link>
      </li>
      <li>
        <router-link to="/history">Question history</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "comp-left"
}
</script>

<style scoped>
.leftmenu {
  padding-top: 150px;
  padding-left: 20px;
  display: inline-block;
  text-align: center;
  position: fixed;
  left: 0;
  z-index: 1000;
}

.admin-sidenav li {
  background-color: peru;
  border: 1px solid #333;
  height: 70px;
  width: 150px;
  float: left;
  margin-top: 5px;
  padding-top: 10px;
}

.admin-sidenav li a {
  text-decoration: none;
  color: #fff;
  font-weight: bold;
  font-size: 17px;
}

.admin-sidenav li a:hover {
  color: rgb(83, 203, 250);
}
</style>
