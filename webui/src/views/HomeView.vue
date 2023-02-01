<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: '',
			response: {},
			error : false
		}
	},
	methods: {
		submit: async function () {
			this.loading = true;
			this.username = this.$refs.username.value;
			if(this.username.length < 3 || this.username.length > 16) {
				this.error = true;
			} else {
				try{
					let identifier = await this.$axios.post("/session", {
						name : this.username
					})
					this.response = identifier.data;
					localStorage.setItem("token", this.response.id);
					localStorage.setItem("name", this.username);
				}catch (e) {
				}
				console.log(this.username + ": " + this.response.id);
				this.loading = false;
				this.$router.push("/homepage");
			}
		},
	},
}
</script>

<template>
	<h1 class="page-title">WASAPhoto</h1>
	<div class = "centering-container">
		<div class="container">
			<h1 class="login-title">Login</h1>
			<form class="login-form" @submit.prevent="submit">
				<label for="username">Username:</label>
				<input v-model="username" type="text" ref="username" placeholder="Insert username..." required>
				<button>Login</button>
			</form>
			<div class="login-error-label" v-show="this.error">
				<h4 class="login-error-text"> The username must be at least 3 characters and a maximum of 16 </h4>
			</div>
		</div>
	</div>
</template>

<style>

.login-error-text {
	color: red;
	text-align: center;
}

.centering-container {
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
}

.page-title{
	position: absolute;
	top: 10%;
	left: 50%;
	transform: translateX(-50%);
}

.container {
	width: 400px;
	height: 400px;
	background-color: #666262;
	display: flex; /* Serve per inserire gli elementi nel container aggiustando la loro dimensione per entrare dentro i bound prima definiti */
	flex-direction: column; /* Direzione in cui vengono piazzati i componenti */
	align-items: center; /* Come vengono allineati i componenti, cioè nel centro rispetto al main axis (column prima definito) */
	justify-content: center; /* Come i componenti vengono piazzati uno rispetto all'altro, in questo caso sono distruibuti centralmente rispetto al main axis */
	border-radius: 5px;
}
 
.login-title{
	text-align: center;
}

.login-form{
	display: flex;
	flex-direction: column;
	align-items: center;
}

.login-form label{
	margin-bottom: 10px;
}

.login-form input{
	width: 100%; /* Occupa tutto lo spazio del componente genitore */
	padding: 12px 20px; /* Specifica lo spazio lasciato a partire dai bordi del componente. In questo caso, top e bottom di 12 e right left di 20 */
	margin: 8px 0; /* Crea spazio ATTORNO ai componenti, quindi top e bottom di 8px e left e right di 0px */
	box-sizing: border-box; /* Include il padding e il maring nel calcolo del width e height */
	border: 1px solid #ccc;
	border-radius: 4px; /* Aggiunge spigoli arrotondati */
}

.login-form button{
	width: 100%;
	background-color: #4caf50;
	color: white;
	padding: 14px 20px;
	margin: 8px 0;
	border: none;
	border-radius: 4px;
	cursor: pointer;
}
</style>
