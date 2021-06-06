<template>
  <div class="right">
    <p style="font-size: 20px; font-weight: bold; text-align: center">Upload Knowledge</p>
    <div class="uploadfile">
      <p style="margin-top: 10px">Upload a file:</p>
      <div>
        <input style="width: 150px" placeholder="Choose a file">
        <button>Upload</button>
      </div>
      <br>
      <p>Or paste it here:</p>
      <textarea style="width: 220px; height: 100px"></textarea>
      <br>
      <button class="btn" v-on:click="myFunction()">Upload</button>
    </div>
    <p id="demo" style="text-align: center; font-size: 20px; font-weight: bold"></p>
  </div>
</template>

<script>
export default {
  name: "CompUpload",
  methods:{
    myFunction: function () {
      document.getElementById("demo").innerHTML = "Upload file success !";
    }
  }
}
</script>

<style scoped>
.right {
  position: absolute;
  margin-top: 150px;
  width: 70%;
  text-align: left;
  margin-left: 15%;
}

.uploadfile {
  margin: 20px;
  text-align: center;
  margin-left: 300px;
  border: 1px solid black;
  background-color: #e5e5e5;
  width: 400px;
  height: 300px;
  font-size: 15px;
}

.uploadfile .btn {
  margin-left: 50px;
  font-weight: bold;
  font-size: 15px;
  background-color: cadetblue;
}
</style>
