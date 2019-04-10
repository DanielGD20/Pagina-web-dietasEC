$(document).ready(function () {

    var contador = 0;

    $("#divInicio").show();
    $("#registroRecetas").hide();
    $("#recetasRegistradas").hide();
    $("#menuRecetas").hide();
    $(".botonBorrar").hide();

    $("#registrar").click(function () {
        $("#divInicio").hide();
        $("#registroRecetas").show();
        $("#recetasRegistradas").hide();
        $("#menuRecetas").hide();
    });

    $("#verRecetas").click(function () {
        $("#divInicio").hide();
        $("#registroRecetas").hide();
        $("#recetasRegistradas").show();
        $("#menuRecetas").hide();
    });

    $("#elaborarMenu").click(function () {
        $("#divInicio").hide();
        $("#registroRecetas").hide();
        $("#recetasRegistradas").hide();
        $("#menuRecetas").show();
    });

    $("#inicio").click(function () {
        $("#divInicio").show();
        $("#registroRecetas").hide();
        $("#recetasRegistradas").hide();
        $("#menuRecetas").hide();
    });


    $(".blanqueador").click(function () {
        contador++;
        console.log(contador);
        $("#grupoDeIngredientes").append(generarTexto(contador));
        if (contador >= 1) {
            $(".botonBorrar").show();
        }
    });

    $(".blanqueador2").click(function () {
        $('#grupo-' + contador).remove();
        contador--;
        if (contador == 0) {
            $(".botonBorrar").hide();
        }
        console.log(contador);
    })
})


function generarTexto(numero) {
    var input = "<div class='input-group' id = 'grupo-" + numero + "'><input type='text' class='form-control' placeholder='Ingrediente' id = 'nombre-" + numero + "'><input type='number' class='form-control' placeholder='Cantidad a utilizar' id = 'cantidad-" + numero + "'></div>";
    console.log(input);
    return input;
}