<template>
  <div class="bottom section-padding">
    <div class="container">
      <div class="row">
        <div class="col-md-12 text-center">
          <div class="copyright">
            <p>© <span>2021</span> <a href="#" class="transition">FPT University</a> All rights reserved.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<!-- Script -->
<script>
export default {
  name: "comp-footer"
}
</script>

<!-- Style -->
<style scoped>

.bottom {
  padding-top: 10px;
  background-color: peru;
  left: 0;
  bottom: 0;
  position: fixed;
  width: 100%;
  height: 60px;
  text-align: center;
  font-weight: bold;
  z-index: 1000;
}

.copyright a {
  color: #f2ff49;
}

.bottom p {
  line-height: 50px;
  font-size: 16px;
}

.bottom .copyright p, .bottom .copyright a:hover {
  color: white;
}

</style>
