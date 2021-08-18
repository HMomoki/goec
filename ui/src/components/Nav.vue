<template>
    <nav class="navbar navbar-light bg-info">
        <ul class="nav d-flex justify-content-start">
            <li class="nav-item">
                <router-link to="/" class="nav-link">Home</router-link>
            </li>
        </ul>
        <ul class="nav d-flex justify-content-end">
            <li class="nav-item">
                <router-link to="/login" class="nav-link">Login</router-link>
            </li>
            <li class="nav-item">
                <router-link to="/register" class="nav-link">Register</router-link>
            </li>
        </ul>
    </nav>
</template>
 
<script>
export default {
  name: "Nav"
}
</script>
 