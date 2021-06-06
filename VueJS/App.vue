<template>
  <div id="app">
    <comp-header-l/>
    <comp-footer/>
    <comp-left/>
    <router-view/>
  </div>
</template>

<script>
import CompHeaderL from "./components/CompHeaderL";
import CompFooter from "./components/CompFooter";
import CompLeft from "./components/CompLeft";


export default {
  name: 'App',
  components: {
    CompHeaderL,
    CompFooter,
    CompLeft
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
