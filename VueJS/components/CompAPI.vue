<template>
  <div class="right">
    <table>
      <tr>
        <th>API name</th>
        <th>Request</th>
      </tr>
      <tr>
        <td>Authenticate</td>
        <td>POST.../api/auth</td>
      </tr>
      <tr>
        <td>Logout</td>
        <td>GET...api/logout</td>
      </tr>
      <tr>
        <td>Ask question</td>
        <td>POST...api/qa</td>
      </tr>
      <tr>
        <td>List of Knowledges</td>
        <td>GET...api/knowledges</td>
      </tr>
      <tr>
        <td>Upload Knowledge</td>
        <td>PUT...api/knowledges</td>
      </tr>
      <tr>
        <td>Question history</td>
        <td>GET...api/history</td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: "CompAPI"
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

table, th, td {
  border: 1px solid black;
  border-collapse: collapse;
  background-color: #e5e5e5;
  margin-left: 200px;
}

th {
  background-color: burlywood;
}

th, td {
  padding: 15px;
  text-align: center;
  width: 200px;
}
</style>
