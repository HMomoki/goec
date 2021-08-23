<template>
    <div class="login">
        <div class="text-center">
            <button @click="sendRequest">リクエスト送信</button>
            <h1>取得結果</h1>
            <p>GET(パラメータ無):<br/><strong>{{title}}</strong></p>
            <p>GET(パラメータ有):<br/><strong>{{name1}}</strong></p>
            <p>POST(form-urlencoded):<br/><strong>{{name2}}</strong></p>
            <p>POST(JSONデータ):<br/><strong>{{company}}</strong></p>
        </div>
    </div>
</template>
 
<script>
import axios from "axios"

export default {
  name: "Login",
  data:() => {
      return{
          title:"",
          name1:"",
          name2:"",
          company:""
        }
    },
    methods: {
        sendRequest: async function(){

            await axios.get("http://localhost:8000/auth/signin").then(
                (response) => {
                    this.title = response.data
                }
            )

            const params = new URLSearchParams();
            params.append("name","青葉");
            await axios.post("http://localhost:8000/auth/signin",params).then(
                (response) => {
                    this.name2 = response.data
                }
            )
        }
    },
    

}
</script>
 
<style>
</style>