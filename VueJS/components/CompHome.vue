<template>
  <div class="right">
    Welcome to < Web name > !<br>
    Our AI can answer all multiple choices questions that may appear in
    a highschool exam.<br>
    First, let <u style="color: dodgerblue"><b><a href="#">LOGIN</a></b></u>. To make an account,
    click <u style="color: dodgerblue"><b><a href="#">REGISTER</a></b></u>. To ask a multiple choices question,
    go to <u style="color: dodgerblue"><b><a href="#">Q&A</a></b></u>. Our AI will give you an answer
    immediately.<br>
    To see all trained knowledges, go to <u style="color: dodgerblue"><b><a href="#">List of
    Knowledges</a></b></u>.<br>
    If you are a teacher, you can use <u style="color: dodgerblue"><b><a href="#">Upload
    Knowledges</a></b></u> to train more
    knowledges
    to the AI by uploading a file.<br>
    To see the history of asked questions, go to <u style="color: dodgerblue"><b><a href="#">Question
    History</a></b></u>.<br>
    Finally, you can use our provied APIs for your program. Go to <u style="color: dodgerblue"><b><a
    href="#">API
    Library</a></b></u> to
    see all available APIs.
  </div>
</template>

<script>
export default {
  name: "CompHomeLogin"
}
</script>

<style scoped>

.right {
  position: absolute;
  margin-top: 150px;
  width: 70%;
  text-align: left;
  margin-left: 20%;
}
</style>
