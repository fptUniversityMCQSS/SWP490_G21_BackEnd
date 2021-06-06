<template>
  <div class="navbar">
    <div class="navbar-left">
      LOGO
    </div>
    <div class="navbar-right">
      <ul>
        <li>
          <router-link to="/home">Home</router-link>
        </li>
        <li><a href="#">API Library</a></li>
        <li><a href="#">About Us</a></li>
        <li>
          <router-link to="/homelogin">Sign in</router-link>
        </li>
        <li><a href="#">Register</a></li>

      </ul>
    </div>
  </div>
</template>

<!-- Script -->
<script>

export default {
  name: 'comp-header'
}
</script>

<!-- Style -->
<style scoped>

.navbar {
  position: fixed;
  top: 0;
  left: 0;
  padding: 20px;
  text-align: center;
  color: white;
  width: 100%;
  font-weight: bold;
  background: peru;
  margin-top: 5px;
}

.navbar .navbar-left {
  float: left;
  margin-left: 30px;
}

.navbar .navbar-right {
  padding-left: 22px;
}

li {
  display: inline-block;
  margin: 0 20px;
}

a {
  color: white;
  text-decoration: none;
}

a:hover {
  color: rgb(83, 203, 250);
}

</style>
