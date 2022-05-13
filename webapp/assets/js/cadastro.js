$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento){
    evento.preventDefault(); // prevenir o comportamento default do form - não recarrega a interface 
    // console.log("to no criar usuario")

    if ( $('#senha').val() != $('#confirmar-senha').val() ) {
        alert("As senhas não correpondem");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){ // statuscode 201 200 204 ... ok deu certo
        alert("Usuário cadastrado com sucesso")
    }).fail(function(erro){ // statuscode 400 404 401 403 500... falhou
        console.log(erro)
        alert("Falha ao cadastrar o usuário")
    });
}