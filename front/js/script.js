$(document).ready(function () {
    // Obtener la fecha actual
  const fechaActual = new Date();
  const dia = fechaActual.getDate().toString().padStart(2, '0');
  const mes = (fechaActual.getMonth() + 1).toString().padStart(2, '0');
  const anio = fechaActual.getFullYear();

 // Asignar la fecha actual al campo de fecha
  $('#fecha').val(`${anio}-${mes}-${dia}`);
    let totalVentasEF = 0; // Variable para el total calculado manualmente (oculta al usuario)
    const nombresChicas = []; // Array para guardar los nombres de las chicas
    const datosTarjetas = []; // Array para guardar los datos de las tarjetas
    const datosEgresos = [];

    // Calcular los totales
    function calcularTotales() {
        // Obtener valores de los campos
        const efectivo = parseFloat($('#efectivo').val()) || 0; // Efectivo total ventas (manual)
        const egresos = parseFloat($('#egresos').val()) || 0;
        const efectivoRendidoCalculado = calcularEfectivoRendido();

        // Calcular total-ventas-EF (oculto)
        totalVentasEF = efectivo - egresos;

        // Actualizar efectivo rendido
        $('#efectivo-rendido').val(efectivoRendidoCalculado.toFixed(2));

        // Calcular totales de las tarjetas
        const { totalTD, totalTC, totalFN } = calcularTotalesTarjetas();
        $('#total-td').val(totalTD.toFixed(2));
        $('#total-tc').val(totalTC.toFixed(2));
        $('#total-fn').val(totalFN.toFixed(2));

        // Calcular total general
        const totalGeneral = efectivoRendidoCalculado + totalTD + totalTC + totalFN;
        $('#total-general').val(totalGeneral.toFixed(2));
    }

    // Calcular efectivo rendido desde la calculadora de billetes
    function calcularEfectivoRendido() {
        let total = 0;

        $('.cantidad').each(function () {
            const denominacion = parseInt($(this).data('denomination')) || 0;
            const cantidad = parseInt($(this).val()) || 0;
            total += denominacion * cantidad;
        });

        return total;
    }
    //jamas en mi vida me toco este error pero es para tapar arreglar el formato de la fecha
    function formatFechaParaServidor(fecha) {
        const fechaObj = new Date(fecha); // Crear un objeto Date a partir de la fecha
        return fechaObj.toISOString(); // Convertir al formato ISO (incluye hora y zona horaria)
    }

    // Calcular totales de las tarjetas
    function calcularTotalesTarjetas() {
        let totalTD = 0; // Total Tarjeta Débito
        let totalTC = 0; // Total Tarjeta Crédito
        let totalFN = 0; // Total Financiera (placeholder, ajustar si es necesario)

        $('#tarjetas-list tr').each(function () {
            const formaPago = $(this).find('.forma-pago').val(); // Forma de pago
            const monto = parseFloat($(this).find('.monto').val()) || 0;

            if (formaPago === 'Débito') {
                totalTD += monto;
            } else if (formaPago === 'Crédito') {
                totalTC += monto;
            } else {
                totalFN += monto; // Otros montos se suman a Financiera
            }
        });

        return { totalTD, totalTC, totalFN };
    }

    // Guardar los nombres de las chicas en un array
    function guardarNombresChicas() {
        // Limpiar el array antes de llenarlo
        nombresChicas.length = 0;

        // Iterar sobre las filas de la tabla y guardar los nombres
        $('#chicas-list tr').each(function () {
            const nombre = $(this).find('.nombre-chica').val().trim();
            if (nombre) {
                nombresChicas.push(nombre); // Añadir al array si no está vacío
            }
        });

        alert(`Nombres guardados: ${nombresChicas.join(', ')}`);
        console.log('Array de chicas:', nombresChicas); // Verificar en la consola
    }

   
   // Guardar los datos de las tarjetas en un array
  
   function guardarDatosTarjetas() {
    datosTarjetas.length = 0; // Limpiar el array antes de llenarlo

    $('#tarjetas-list tr').each(function () {
        const tipo = $(this).find('.forma-pago').val(); // Convertir a entero
        const local = parseInt($('#sucursal').val(), 10); // Convertir a entero
        const tarjetaFinanciera = $(this).find('.tarjeta-financiera').val(); // Valor como string
        const cuotas = parseInt($(this).find('.cuotas').val(), 10) || 0; // Convertir a entero o 0
        const monto = parseFloat($(this).find('.monto').val()) || 0; // Convertir a flotante o 0

        // Validar que tipo es un entero y los demás valores son válidos
        if (local && tarjetaFinanciera && monto > 0) {
            datosTarjetas.push({
                local,
                tipo,
                tarjetaFinanciera,
                cuotas, 
                monto,
            });
        } else {
            // Mostrar alerta si algún dato no es válido
            alert("NO es int capo");
        }
    });

    console.log('Array de tarjetas:', datosTarjetas); // Verificar en la consola

    function guardarDatosEgresos() {
        datosEgresos.length = 0; // Limpiar el array antes de llenarlo
    
        $('#egresos-list tr').each(function () {
            const concepto = $(this).find('.concepto').val().trim();
            const monto = parseFloat($(this).find('.monto').val()) || 0;
    
            if (concepto && monto > 0) {
                datosEgresos.push({
                    concepto: concepto,
                    monto: monto,
                });
            }
        });
    
        //console.log('Array de egresos:', datosEgresos); // Verificar en la consola
    }}

    function guardarDatosEgresos() {
        datosEgresos.length = 0; // Limpiar el array antes de llenarlo
    
        $('#egresos-list tr').each(function () {
            const concepto = $(this).find('.concepto').val().trim();
            const monto = parseFloat($(this).find('.monto').val()) || 0;
    
            if (concepto && monto > 0) {
                datosEgresos.push({
                    concepto: concepto,
                    monto: monto,
                });
            }
        });
    
}   

    // configurar el boton con la funcion
    $('#guardar').click(confirmarSubida)
    // Validar y confirmar subida de datos a la base de datos
    function confirmarSubida() {
    // Calcular efectivo rendido y diferencia
    const efectivoRendidoCalculado = parseFloat($('#efectivo-rendido').val()) || 0;
    const diferencia = (totalVentasEF - efectivoRendidoCalculado).toFixed(2);

    // Verificar si existe diferencia
    if (diferencia != 0) {
        const mensaje = `
        Existe una discrepancia entre el total-ventas-EF y el efectivo rendido:
        - Total Ventas EF (calculado): ${totalVentasEF.toFixed(2)}
        - Efectivo Rendido: ${efectivoRendidoCalculado.toFixed(2)}
        - Diferencia: ${diferencia}

        ¿Estás seguro de que quieres subir los datos a la base de datos a pesar de esta diferencia?
        `;
        if (!confirm(mensaje)) {
            return; // Detener si el usuario no confirma
        }
    }

    // Proceder con la subida de datos
    console.log('Nombres de chicas:', nombresChicas); // Mostrar nombres de chicas en la consola
    console.log('Datos de tarjetas:', datosTarjetas); // Mostrar datos de tarjetas en la consola
    alert('Datos confirmados y listos para enviar a la base de datos.');
    guardarNombresChicas(); // Guardar los nombres antes de subir
    guardarDatosTarjetas(); // Guardar los datos de las tarjetas antes de subir
    guardarDatosEgresos();
    subirDatos(); // Llamar a la función que sube los datos
 }

// Subir datos a la base de datos 
  function subirDatos() {
    console.log("ID Local seleccionado:", $('#sucursal').val());
    console.log("ID Local seleccionado:", $('#usuario-id').val());
    const turnoData = {
        id_local: parseInt($('#sucursal').val(), 10), // ID de la sucursal
        fecha: formatFechaParaServidor($('#fecha').val()), // Fecha formateada
        tipo: $('#turno').val(), // Tipo de turno (Mañana/Tarde)
        ingreso_ef: parseFloat($('#efectivo').val()) || 0, // Efectivo
        ingreso_debito: parseFloat($('#total-td').val()) || 0, // Tarjeta débito
        ingreso_credito: parseFloat($('#total-tc').val()) || 0, // Tarjeta crédito
        financiera: parseFloat($('#total-fn').val()) || 0, // Financiera
        egreso_total: parseFloat($('#egresos').val()) || 0, // Total de egresos
        diferencia_caja: parseFloat($('#diferencia-caja').val()) || 0, // Diferencia de caja
        id_usuario: parseInt($('#usuario-id').val(), 10), // ID del usuario
    };

    // Crear el array de datos de los vouchers
    const vouchersData = datosTarjetas.map(tarjeta => ({
        tipo: tarjeta.tipo, // "Crédito" o "Débito"
        tarjeta: tarjeta.tarjetaFinanciera, // Nombre de la tarjeta
        cuotas: parseInt(tarjeta.cuotas, 10) || 0, // Número de cuotas
        monto: parseFloat(tarjeta.monto) || 0, // Monto del voucher
        id_turno: turnoData.id_local, // ID del turno relacionado
    }));

    
    // Crear el array de datos de los egresos
    const conceptosEgresoData = datosEgresos.map(egreso => ({
        concepto: egreso.concepto,
        monto: egreso.monto,
        id_turno: turnoData.id_local,
    }));
    console.log('Array de egresos lo que se deberia enviar:', conceptosEgresoData);
   

    // Crear el array de datos de las chicas
    const chicasData = nombresChicas.map(nombre => ({
        Nombre: nombre // Nombre de la chica
    }));

    // Crear el objeto final para enviar
    const dataToSend = {
        turno: turnoData,
        vouchers: vouchersData,
        ConceptoEgreso: conceptosEgresoData,
        chicas: chicasData,
    };
    
    // Enviar datos a través de una solicitud POST
    $.ajax({
        url: '/cargar', // Cambia esta URL a la ruta correcta en tu servidor
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(dataToSend),
        success: function (response) {
            alert('Datos enviados con éxito');
            console.log(response);
        },
        error: function (error) {
            alert('Error al enviar los datos');
            console.error(error);
        }
    });
  }
    // Escuchar cambios en los campos relevantes
    $('#efectivo, #egresos, .cantidad').on('input', calcularTotales); // Actualizar totales cuando cambian los billetes o valores principales
    $(document).on('input', '.monto, .forma-pago', calcularTotales); // Actualizar totales cuando cambian las filas de tarjetas

    // Evento para el botón "Calcular Totales"
    $('#calcular-totales').click(calcularTotales);
    /*Agregar renglon para chica*/ 

    $('#agregar-chica').click(function(){
        const nuevaFila = `
        <tr>
         <td><input type="text" placeholder="Nombre" class="nombre-chica"></td>
        </tr>`;
        $('#chicas-list').append(nuevaFila);
         
    })
    $(document).ready(function () {
         
  
// Función para guardar los datos de los egresos en un array

        // Función para calcular el total de egresos
    
    function calcularTotalEgresos() {
        let totalEgresos = 0;
    
        $('#egresos-list tr').each(function () {
            const monto = parseFloat($(this).find('.monto').val()) || 0;
            totalEgresos += monto;
        });
    
         $('#egresos').val(totalEgresos.toFixed(2)); // Actualizar el campo de egresos
       
        }
    
        // Evento para agregar filas en la tabla de egresos
        $('#agregar-egreso').click(function () {
            const nuevaFila =
                `<tr>
                    <td><input type="text" placeholder="Concepto" class="concepto"></td>
                    <td><input type="number" placeholder="Monto" class="monto"></td>
                </tr>`;
            $('#egresos-list').append(nuevaFila);
        });
    
        // Evento para guardar los datos de los egresos y calcular el total
        $('#guardar').click(function () {
            guardarDatosEgresos(); // Guardar los datos en el array
            calcularTotalEgresos(); // Calcular el total de egresos
        });
    
        // Escuchar cambios en los montos de la tabla de egresos
        $(document).on('input', '.monto', calcularTotalEgresos);
    });

    // Evento para agregar filas en la tabla de tarjetas
    $('#agregar-fila').click(function () {
        const nuevaFila = `
                         <tr>
                            <td>
                                <!-- Select para el tipo de pago -->
                                <select class="forma-pago">
                                    <option value="">Seleccionar</option>
                                    <option value="Débito">Débito</option>
                                    <option value="Crédito">Crédito</option>
                                </select>
                            </td>
                            <td>
                                <!-- Select para las tarjetas -->
                                <select class="tarjeta-financiera">
                                    <option value="">Seleccionar</option>
                                    {{ range .tarjetas }}
                                        <option value="{{ .Nombre }}" data-tipo="{{ .Tipo }}">{{ .Nombre }}</option>
                                    {{ end }}
                                </select>
                            </td>
                            <td><input type="number" placeholder="Cuotas" class="cuotas"></td>
                            <td><input type="number" placeholder="Monto" class="monto"></td>
                        </tr>`;
        $('#tarjetas-list').append(nuevaFila); // Agregar fila a la tabla
    });

    $(document).ready(function () {
        // Delegar el evento "change" en el contenedor
        $('#tarjetas-list').on('change', '.forma-pago', function () {
            const tipoSeleccionado = $(this).val(); // Obtener el tipo seleccionado (Débito o Crédito)
            const $fila = $(this).closest('tr'); // Obtener la fila actual
            const $selectTarjetas = $fila.find('.tarjeta-financiera'); // Obtener el select de tarjetas
    
            // Limpiar las opciones actuales
            $selectTarjetas.html('<option value="">Seleccionar</option>');
    
            // Recuperar las opciones originales de las tarjetas
            const tarjetasOriginales = $('#tarjetas-list').data('tarjetas'); // Almacenar las tarjetas una vez al cargar la página
    
            // Filtrar y agregar las tarjetas correspondientes
            tarjetasOriginales.forEach(tarjeta => {
                if (tarjeta.tipo === tipoSeleccionado) {
                    const opcion = `<option value="${tarjeta.nombre}" data-tipo="${tarjeta.tipo}">${tarjeta.nombre}</option>`;
                    $selectTarjetas.append(opcion);
                }
            });
        });
    
        // Cargar las tarjetas originales desde el backend al iniciar
        const tarjetas = [];
        $('.tarjeta-financiera option[data-tipo]').each(function () {
            tarjetas.push({
                nombre: $(this).val(),
                tipo: $(this).data('tipo')
            });
        });
    
        // Almacenar las tarjetas originales en el contenedor
        $('#tarjetas-list').data('tarjetas', tarjetas);
    });
    

});

