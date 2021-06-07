<template>
  <div class="right">
    <table>
      <tr>
        <th>Title</th>
        <th>Asked date</th>
        <th>Answer</th>
      </tr>
      <tr>
        <td>Jordan Ragsdale has been the chair of a committee</td>
        <td>21/05/2021 19:00</td>
        <td>A</td>
      </tr>
      <tr>
        <td>Jordan Ragsdale has been the chair of a committee</td>
        <td>20/05/2021 20:00</td>
        <td>B</td>
      </tr>
      <tr>
        <td>Jordan Ragsdale has been the chair of a committee</td>
        <td>19/05/2021 17:00</td>
        <td>A</td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: "CompHistory"
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
