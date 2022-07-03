<template>
    <div class="resister">
      <div class="container py-5">
        <div class="d-flex flex-column flex-lg-row justify-content-center">
          <div class="col-12 col-lg-6 p-3 p-sm-5 mr-lg-1 mb-3 mb-lg-0" style="background: #efefef">
            <div class="text pb-3">会員登録</div>
            <form @submit.prevent="register">
              <input v-model="email" type="email" class="form-control" placeholder="email">
              <input v-model="password" type="password" class="form-control" placeholder="Password">
              <button class="w-100 btn btn-lg btn-primary" type="submit">登録</button>
            </form>
          </div>
        </div>
      </div>
    </div>
</template>
 
<script>
import axios from "axios"

axios.defaults.baseURL = 'http://localhost:2500'

export default {
  name: "Register",
  methods: {
    register : async function(){
      const params = new URLSearchParams();

      params.append("email",this.email);
      params.append("password",this.password);

      await axios.post("/auth/signup",params)
    }
  },
}
</script>
 
<style>
</style>