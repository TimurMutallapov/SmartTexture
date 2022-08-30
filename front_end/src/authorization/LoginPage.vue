<template>
  <v-app>
    <v-card width="400" class="mx-auto mt-5">
      <v-card-title class="pb-0">
        <h4>Login</h4>
      </v-card-title>
      <v-card-text>
        <v-form v-model="valid">
          <v-text-field 
            label="Username" 
            prepend-icon="mdi-account-circle"
            v-model="data.email"
            id="email" 
          />
          <v-text-field 
            :type="showPassword ? 'text' : 'password'" 
            label="Password"
            id="password" 
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
            v-model="data.password"
          />
        </v-form>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions @submit.prevent="">
        <v-btn  to= '/signup' color="success">Register</v-btn>
        <v-btn id="myForm" v-on:click="submit" color="info"  >Login</v-btn>
      </v-card-actions>
    </v-card>
  </v-app>
</template>

<script> 


import { reactive } from "vue"
import { useRouter } from "vue-router"


export default { 
  el: '#root',
  name: 'App', 
  data () { 
    return { 
      showPassword: false 
   } 
 }, 
  setup() {
    const data = reactive({
      email: "",
      password: "",
    })
    const router = useRouter()
    return {
      data,
      router,
    }
  },
   methods: {
    async submit() {
      
      const url = "http://localhost:5000/auth/login"
      const res = await fetch(url,  {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        credentials: "include",
        body: JSON.stringify(this.data)
      })
      const status = res.status
      if (status == 200) {
        
        this.router.push("/")
      } else {
        console.log("no auth")
      }
    }
  }
}

</script>