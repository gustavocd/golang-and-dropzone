(function (Dropzone) {
    Dropzone.autoDiscover = false;
    
    let drop = new Dropzone('div#filesBox', {
        maxFiles: 5,
        url: '/upload',
        method: 'post',
        addRemoveLinks: true,
        acceptedFiles: 'image/*',
        createImageThumbnails: true,
        dictRemoveFile: 'Quitar archivo',
        dictMaxFilesExceeded: 'No puedes subir más de 5 imágenes',
        dictDefaultMessage: 'Haz click o arrastra aquí tus imágenes',
        dictInvalidFileType: 'Solo puedes subir archivos .png, .jpg  y .svg'
    });

    drop.on('success', function(file, response) {
        file.id = response.id
    });

    drop.on('removedfile', function(file) {
        console.log(file);
    });
})(Dropzone);