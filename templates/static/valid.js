 //validation code for input
 var user=document.forms['form']['user'];
 var password=document.forms['form']['password'];

 var username_error=document.getElementById('username_error');
 var pass_error=document.getElementById('pass_error');

 function validated(){
 	if(user.value.length < 10){
 		user.stype.border="3px solid red";
 		username_error.style.display="block";
 		user.focus();
 		return false;
 	}
 }