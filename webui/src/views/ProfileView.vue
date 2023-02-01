<script>
export default {
    data() {
        return {
            profileName : "",
            profileData : [],
            followButtonLabel : "Follow",
            banButtonLabel : "Ban",
            followerNumber : 0,
            banColor : {r : 199, g: 29, b: 29, a: 1},
            followColor : {r : 41, g: 216, b: 47, a: 1},
            personalProfile: false,
            userNotFound: false
        }
    },
    async mounted() {
        try{   
            this.profileName = localStorage.getItem("profile")
            this.personalProfile = this.profileName == localStorage.getItem("name")
            if(this.personalProfile){
                this.$router.replace("/homepage/profile/personal")
            }
            let response = await this.$axios.get("/profiles/" + this.profileName)
            this.profileData = response.data
            console.log(this.profileData)
            if (this.profileData["Follower"] != null) {
                this.followerNumber = this.profileData["Follower"] == null ? 0 : this.profileData["Follower"].length
                for(let i = 0; i < this.profileData["Follower"].length; i++){
                    if(localStorage.getItem("name") == this.profileData["Follower"][i]["name"]){
                        this.followButtonLabel = "Unfollow"
                        break
                    }
                }
            }
            let risp = await this.$axios.get("/profiles/" + localStorage.getItem("name") + "/banned/" + localStorage.getItem("profile"))
            let dataBan = risp.data["banned"]
            this.banButtonLabel = dataBan == "true" ? "Unban" : "Ban"

            if(this.banButtonLabel == "Unban") {
                this.followColor.a = 0.2
            }
            if(this.followButtonLabel == "Unfollow") {
                this.banColor.a = 0.2
            }
        } catch (e) {
            this.userNotFound = true;
        }
    },
    methods : {
        async follow() {
            if (this.banButtonLabel != "Unban"){
                if (this.followButtonLabel == "Follow") {
                    this.followerNumber = this.followerNumber + 1
                    this.followButtonLabel = "Unfollow"
                    await this.$axios.put("/profiles/" + localStorage.getItem("name") + "/followers/" + this.profileName, {
                        name : this.profileName
                    })
                    this.banColor.a = 0.2
                } else {
                    this.followerNumber = this.followerNumber - 1
                    this.followButtonLabel = "Follow"
                    await this.$axios.delete("/profiles/" + localStorage.getItem("name") + "/followers/" + this.profileName)
                    this.banColor.a = 1
                }
            }
        },
        async ban() {
            if (this.followButtonLabel != "Unfollow"){
                if (this.banButtonLabel == "Ban") {
                    this.banButtonLabel = "Unban"
                    await this.$axios.put("/profiles/" + localStorage.getItem("name") + "/banned/" + localStorage.getItem("profile"), {
                        name : localStorage.getItem("profile")
                    })
                    this.followColor.a = 0.2
                } else {
                    this.banButtonLabel = "Ban"
                    await this.$axios.delete("/profiles/" + localStorage.getItem("name") + "/banned/" + localStorage.getItem("profile"))
                    this.followColor.a = 1
                }
            }
            this.update();
        },
        update() {
            this.$router.go(0);
        }
    }
}
</script>

<template>
    <div class="page">
        <Header @updateParent="update"></Header>
        <div class="user-not-found" v-show="this.userNotFound">
            <h1 class="user-not-found-label"> User Not Found</h1>
        </div>
        <div class="profile-info" v-show="!this.userNotFound">
            <div class="profile-name"> {{ this.profileName }} </div>

            <div class="general-button">
                <button class="ban-button" 
                    @click="ban" 
                    :style="{ backgroundColor : 'rgba(' + banColor.r + ', ' + banColor.g + ', ' + banColor.b + ', ' + banColor.a + ')'}">
                    {{ this.banButtonLabel }}
                </button>
                <button class="follow-button" 
                    @click="follow"
                    :style="{ backgroundColor : 'rgba(' + followColor.r + ', ' + followColor.g + ', ' + followColor.b + ', ' + followColor.a + ')'}"> 
                    {{ this.followButtonLabel }}
                </button>
            </div>

            <div class="profile-info-box" v-show="this.banButtonLabel !== 'Unban'">
                <div class="profile-follower-box">
                    <div class="profile-follower"> Follower </div>
                    <div class="profile-follower-number"> {{ this.followerNumber }} </div>
                </div>

                <div class="profile-following-box">
                    <div class="profile-following"> Following </div>
                    <div class="profile-following-number"> {{ this.profileData["Following"] == null ? 0 : this.profileData["Following"].length }} </div>
                </div>

                <div class="profile-photoNumber-box">
                    <div class="profile-photoNumber"> Photo </div>
                    <div class="profile-photoNumber-number"> {{ this.profileData["PhotoNumber"] == null ? 0 : this.profileData["PhotoNumber"] }} </div>
                </div>
            </div>
        </div>
        <div class="photo-container" v-show="this.banButtonLabel !== 'Unban'">
            <div class="photo" v-if="this.profileData['Photos'] != null" v-for="photo in this.profileData['Photos']" :key="photo.photoID">
                <Photo :id="photo.photoID" :name="this.profileName" :likes="photo.likeNumber" :comments="photo.commentNumber" :date="photo.date" :textComments="photo.comments"/>
            </div>
        </div>
    </div>
</template>

<style>

.user-not-found-label {
    font-weight: bold;
    font-size: 20px;
    color: red;
}

.ban-button {
    width: 100px;
    height: 30px;
    position: relative;
    bottom: 45px;
    left: 810px;
	color: white;
	border: none;
	border-radius: 4px;
}

.follow-button {
    width: 100px;
    height: 30px;
    position: relative;
    bottom: 45px;
    left: 600px;
	color: white;
	border: none;
	border-radius: 4px;
}

.photo-container {
    margin-top: 20px;
}

.profile-photoNumber {
    font-weight: bold;
    font-size: 30px;
}

.profile-photoNumber-number {
    font-weight: bold;
    font-size: 30px;
}

.profile-following {
    font-weight: bold;
    font-size: 30px;
}

.profile-following-number {
    font-weight: bold;
    font-size: 30px;
}

.profile-follower {
    font-weight: bold;
    font-size: 30px;
}

.profile-follower-number {
    font-weight: bold;
    font-size: 30px;
}

.profile-photoNumber-box {
    width: 33.33%;
}

.profile-following-box {
    width: 33.33%;
}

.profile-follower-box {
    width: 33.33%;
}

.profile-info-box {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    text-align: center;
    position: relative;
    bottom: 10px;
}

.profile-name {
    font-weight: bold;
    font-size: 40px;
    margin-left: 10px;
}

.profile-info {
    border: 1px solid #ccc;
    border-radius: 4px;
    width: 50%;
    height: 200px;
}

.page {
    display: flex;
    flex-direction: column;
    align-items: center;
}

</style>