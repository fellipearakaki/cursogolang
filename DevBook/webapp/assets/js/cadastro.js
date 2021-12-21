$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(){
    evento.preventDefault();
    console.log("Dentro da funçao usuario")
}

$.ajax({
    url:"/criar-usuario",
    method: "POST",
    data:{
        nome: $("#nome").val(),
        cpf: $("#cpf").val(),

    }

})