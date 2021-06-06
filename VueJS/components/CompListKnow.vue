<template>
  <div class="right">
    <table>
      <tr>
        <th>File name</th>
        <th>Upload date</th>
      </tr>
      <tr>
        <td>Fullstacks-Web-Dev-Resource.pdf</td>
        <td>21/05/2021 19:00</td>
      </tr>
      <tr>
        <td>Fullstacks-Web-Dev-Resource.pdf</td>
        <td>21/05/2021 19:00</td>
      </tr>
      <tr>
        <td>System_Analysis_Design.pdf</td>
        <td>20/05/2021 22:35</td>
      </tr>
      <tr>
        <td>Computer_Architecture.pdf</td>
        <td>19/05/2021 16:54</td>
      </tr>
      <tr>
        <td>B_Giao_dc_va_dao_tao.pdf</td>
        <td>18/01/2021 11:22</td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: "CompListKnow"
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
