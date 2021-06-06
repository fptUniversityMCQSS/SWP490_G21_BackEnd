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
        <li>
          <router-link to="/api">API Library</router-link>
        </li>
        <li><a href="#">About Us</a></li>
        <li>
          <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
        </li>
        <li>
          <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
        </li>
        <li>
          <div>
            <b-dropdown text="Profile">
              <b-dropdown-item>Account</b-dropdown-item>
              <b-dropdown-item>Logout</b-dropdown-item>
            </b-dropdown>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "CompHeaderLogin"
}
</script>

<style scoped>

.navbar {
  position: fixed;
  top: 0;
  left: 0;
  padding: 15px;
  text-align: center;
  color: white;
  width: 100%;
  font-weight: bold;
  background: peru;
  z-index: 1000;
  height: 70px;
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
