$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();//não da refresh na tela

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não coincidem!")
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
    });
}