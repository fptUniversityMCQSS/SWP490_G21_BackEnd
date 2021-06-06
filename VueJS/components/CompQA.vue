<template>
  <div class="right">
    <p style="font-size: 20px; font-weight: bold; text-align: center">Question and Answer</p>
    <ul class="tabs">
      <li v-on:click="changeTabs('one')">
        <b-col lg="4" class="pb-2">
          <b-button style="width: 100px; background-color: peru">Option 1</b-button>
        </b-col>
      </li>
      <li v-on:click="changeTabs('two')">
        <b-col lg="4" class="pb-2">
          <b-button style="width: 100px; background-color: peru">Option 2</b-button>
        </b-col>
      </li>
    </ul>

    <div class="doc">
      <p id="doc" style="text-align: center; font-size: 20px; font-weight: bold"></p>
    </div>

    <div class="option1" v-if="tabSelected === 'one'">
      <p style="margin-top: 10px">Ask your question:</p>
      <textarea style="width: 220px; height: 60px"></textarea><br>
      <button class="btn" v-on:click="myFunction()">Send</button>
    </div>

    <div class="option2" v-else>
      <p style="margin-top: 10px">Ask your question:</p>
      <div>
        <div><textarea style="width: 200px; height: 60px"></textarea></div>
        <div>
          <b-form-spinbutton id="sb-inline" v-model="value" min="1" max="3" inline></b-form-spinbutton>
        </div>
      </div>
      <br>
      <div v-if="value === 1">
        <input placeholder="Answer"><br>
      </div>
      <div v-else-if="value === 2">
        A <input placeholder="Answer A"><br>
        B <input placeholder="Answer B"><br>
      </div>
      <div v-else-if="value === 3">
        A <input placeholder="Answer A"><br>
        B <input placeholder="Answer B"><br>
        C <input placeholder="Answer C"><br>
      </div>
      <br>
      <button class="btn" v-on:click="myFunction()">Send</button>
    </div>

    <p id="demo" style="text-align: center; font-size: 20px; font-weight: bold"></p>
  </div>

</template>

<script>
export default {
  name: "CompQA",
  data() {
    return {
      tabSelected: "one",
      value: 1
    }
  },
  methods: {
    changeTabs(tab) {
      this.tabSelected = tab;
    },
    myFunction: function () {
      document.getElementById("demo").innerHTML = "Answer !";
      document.getElementById("doc").innerHTML = "Document about answer !";
    }
  }
}
</script>

<style scoped>
.right {
  position: absolute;
  margin-top: 100px;
  width: 80%;
  text-align: left;
  margin-left: 10%;
}

.doc {
  position: absolute;
  margin-left: 800px;
}

.option1 {
  margin: 20px;
  text-align: center;
  margin-left: 300px;
  border: 1px solid black;
  background-color: #e5e5e5;
  width: 400px;
  height: 170px;
  font-size: 15px;
}

.option1 .btn {
  font-weight: bold;
  font-size: 15px;
  background-color: peru;
}

.option2 {
  margin: 20px;
  text-align: center;
  margin-left: 300px;
  border: 1px solid black;
  background-color: #e5e5e5;
  width: 400px;
  height: auto;
  font-size: 15px;
}

.option2 .btn {
  font-weight: bold;
  font-size: 15px;
  background-color: peru;
}

.tabs {
  margin-left: 30%;
}

.tabs li {
  display: inline-block;
  padding-left: 40px;
}
</style>
