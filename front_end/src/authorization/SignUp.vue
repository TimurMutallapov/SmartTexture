<template>
  <v-app>
    <v-card width="400" class="mx-auto mt-5">
      <v-card-title class="pb-0">
        <h4>Registration</h4>
      </v-card-title>
      <v-card-text>
        <v-form v-model="valid">
            
          <v-text-field 
            label="Enter you email" 
            v-model.lazy="form.email"
            :rules="emailRules"
            id="email" 
            color="blue-grey "
            prepend-icon="mdi-account-circle"
            required
          />
          <v-text-field 
            label="Enter you name" 
            v-model.lazy="form.name"
            :rules="rulName"
            :counter="10"
            color="teal"
            prepend-icon="mdi-account-circle"
            required
          />
          <v-text-field 
            :type="showPassword ? 'text' : 'password'" 
            label="Retern password"
            v-model.lazy="form.password"
            counter="25"
            maxlength="25"
            color="red darken-1"
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
            required
          />
          <v-text-field 
            :type="showPassword ? 'text' : 'password'" 
            label="Repeat password"
            v-model="form.confirm"
            counter="25"
            color="red darken-1"
            prepend-icon="mdi-lock"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
          />
            
        </v-form>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-btn @click="sumbitForm" variant="outlined" color="red">SignUP</v-btn>
      </v-card-actions>
    </v-card>
  </v-app>
</template>

<script> 
import useValidate from "@vuelidate/core"
import {required, email, sameAs} from '@vuelidate/validators'
import { reactive } from "vue"
import { useRouter } from "vue-router"


export default { 
  name: 'App', 
  data: () => { 
      return { 
      valid: false,
      v$: useValidate(),
        
      emailRules: [
        v => /.+@.+/.test(v) || 'E-mail must be valid'
      ],
      rulName: [
         v => !!v || 'Name is required',
         v => v.length <= 10 || 'Name must be less than 10 characters'
      ],
      showPassword: false ,
          } 
  } ,
    validations (){
    return{
      form:{
        email: {required, email},
        password: {required},
        confirm: { required, sameAsPassword: sameAs(this.form.password) }
      //confirm: {required}
      }
      
    }
  },
  setup() {
    const form = reactive({
      email: "",
      password: "",
      name:''
    })
    const router = useRouter()
    return {
      form,
      router,
    }
  },
  methods:{
   async sumbitForm() {
      this.v$.$validate()
      if (!this.v$.$error) {
        const url = "http://localhost:5000/auth/register"
        const res = await fetch(url,  {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        credentials: "include",
        body: JSON.stringify(this.form)
      })
      const status = res.status
      if (status == 200) {
        
        this.router.push("/login")
      } else {
        alert('server error')
      }
        
         } else { 
          alert('incorrect username or password, or u use temp mail')
      }
     },
  },
} 
</script>