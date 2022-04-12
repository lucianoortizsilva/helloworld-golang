$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert('Ops...As senhas não coincidem!')
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
    }).done(function() { // 201 200 204 ..
        alert('Usuário cadastrado com sucesso!')
    }).fail(function(erro) { // 400 404 401 403 500 ...
        console.log(erro);                    
        alert('Erro ao cadastrar usuário!')
    });    
}