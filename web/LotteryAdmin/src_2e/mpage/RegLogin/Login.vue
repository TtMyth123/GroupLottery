<template>
  <div class="page_login">
    <!-- 头部logo -->
    <div class="head">
      <div class="head_bg">
        <div class="head_inner_bg">
          <img style="width: 55px;height: 65px;" src="./images/head.png" class="head_logo" />
        </div>
      </div>
    </div>
    <div class="title_box">
      <span class="title">
      Hệ thống quản lý
      </span>

    </div>

    <!-- 登录form -->
    <div class="login_form">
      <div class="input">
        <div class="img">
          <img style="width:27px;height: 27px;" src="./images/icon_user.png" />
        </div>
        <input type="text" v-model="username" placeholder="请输入用户账号">
        <div class="img">
          <img @click="delUser" class="img_del" src="./images/icon_del.png" />
        </div>
      </div>
      <div class="line" />
      <div class="input">
        <div class="img">
          <img style="width:20px;height: 25px;" src="./images/icon_pwd.png" />
        </div>
        <input :type="pwdType" :value="userpwd" @input="inputPwd" placeholder="请输入密码">
        <div class="img" @click="switchPwd">
          <img class="img_pwd_switch" src="./images/icon_pwd_switch.png" />
        </div>
      </div>
    </div>
    <!-- 登录提交 -->
    <div class="login_btn">
      <button class="submit" type="primary" @click="login">登录</button>
    </div>
    <div class="quick_login_line">
      <span class="line"> </span>
    </div>
    <div class="opts">
      <span @click="goReg" class="text">立即注册</span>
      <span @click="findPwd" class="text">忘记密码？</span>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      username: '',
      userpwd: '',
      pwdType: 'password',
    }
  },
  methods: {
    inputUsername(e) {
      this.username = e.target.value
    },
    inputPwd(e) {
      this.userpwd = e.target.value
    },
    delUser() {
      this.username = ''
    },
    switchPwd() {
      this.pwdType = this.pwdType === 'text' ? 'password' : 'text'
    },
    login() {
      console.log('username:' + this.username + ',pwd:' + this.userpwd)
    },
    findPwd() {
    },
    goReg() {
    },
    thirdLogin(type) {
      console.log(type)
    }
  }
}
</script>
<style lang="scss" scoped>
$logo-padding: 20px;
$form-border-color: rgba(214, 214, 214, 1);
$text-color: #B6B6B6;

page {
  height: auto;
  min-height: 100%;
  background-color: #f5f6f8;
}

.page_login {
  padding: 10px;
}

.head {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: $logo-padding;
  padding-bottom: $logo-padding;

  .head_bg {
    border: 1px solid #fbecf1;
    border-radius: 50px;
    width: 100px;
    height: 100px;
    display: flex;
    align-items: center;
    justify-content: center;

    .head_inner_bg {
      border-radius: 40px;
      width: 80px;
      height: 80px;
      display: flex;
      background-color: #fc2c5d;
      align-items: flex-end;
      justify-content: center;
      overflow: hidden;
    }
  }
}
.title_box {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.title {

}

.login_form {
  display: flex;
  margin: 20px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 1px solid $form-border-color;
  border-radius: 10px;

  .line {
    width: 100%;
    height: 1px;
    background-color: $form-border-color;
  }

  .input {
    width: 100%;
    max-height: 45px;
    display: flex;
    padding: 3px;
    flex-direction: row;
    align-items: center;
    justify-content: center;

    .img {
      min-width: 40px;
      min-height: 40px;
      margin: 5px;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .img_del {
      width: 21px;
      height: 21px;
    }

    .img_pwd_switch {
      width: 28px;
      height: 12px;
    }

    input {
      outline: none;
      height: 30px;
      width: 100%;

      &:focus {
        outline: none;
      }
    }
  }
}

.submit {
  margin-top: 10px;
  margin-left: 20px;
  margin-right: 20px;
  padding: 8px 40px;
  border-radius: 8px;
  border: 0px solid;
  color: white;
  background-color: rgba(252, 44, 93, 1.0);
  -webkit-tap-highlight-color: rgba(252, 44, 93, 1.0);

  &:active {
    color: #B6B6B6;
    background-color: rgba(252, 44, 93, 0.8);
  }
}

.login_btn {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.opts {
  margin-top: 5px;
  margin-left: 25px;
  margin-right: 25px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  .text {
    font-size: 13px;
    color: $text-color;
  }
}

.quick_login_line {
  margin-top: 40px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;

  .line {
    height: 1px;
    width: 100%;
    background-color: rgba(211, 211, 213, 1);
  }

  .text {
    font-size: 13px;
    color: rgba(200, 200, 200, 1);
    margin: 2px;
  }
}

.quick_login_list {
  margin-top: 30px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;

  .item {
    width: 50px;
    height: 50px;
    margin: 20px;
  }
}
</style>
