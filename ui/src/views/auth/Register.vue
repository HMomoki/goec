<template>
  <div class="container py-5">
    <h1 class="heading--medium text-center mb-4 mb-sm-5">会員登録</h1>
    <div class="d-flex flex-column flex-lg-row justify-content-center">
      <div class="col-12 col-lg-6 p-3 p-sm-5 mr-lg-1 mb-3 mb-lg-0 bg-light">
        <div class="text pb-3">基本情報入力</div>
        <form @submit.prevent="submit">
          <div class="form-group">
            <label for="email" class="font-weight-bold mb-1">メールアドレス</label>
            <input v-model="email" type="email" class="form-control border-0 mb-3" placeholder="Email" required>
          </div>
          <div class="form-group">
            <label for="password" class="font-weight-bold mb-1">パスワード</label>
            <input v-model="password" type="password" class="form-control border-0 mb-3" placeholder="Password" required>
          </div>
          <div class="text-center">
            <button class="btn btn-lg btn-primary col-12 col-sm-6 rounded-0 mt-4" type="submit">登録</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
 
<script>
export default {
  name: "Register"
}
</script>
 
<style>
</style>