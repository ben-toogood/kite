import { useState } from "react";
import { gql, useMutation } from "@apollo/client";
import { createPost } from "./gql-types";

const validateFile = (file: File) => {
  const validTypes = [
    "image/jpeg",
    "image/jpg",
    "image/png",
    "image/gif",
    "image/x-icon",
  ];
  if (validTypes.indexOf(file.type) === -1) {
    return false;
  }
  return true;
};

const POSTUPLOAD = gql`
  mutation createPost($description: String!, $file: Upload!) {
    createPost(description: $description, image: $file) {
      imageURL
    }
  }
`;

const Upload = () => {
  const [previewImage, setPreviewImage] = useState<string | undefined>(
    undefined
  );
  const [selectedFile, setSelectedFile] = useState<File | undefined>(undefined);
  const [description, setDescription] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [uploadFile, { loading, data }] = useMutation<createPost>(POSTUPLOAD);

  const handleFiles = (files: FileList) => {
    if (files.length > 1) {
      setErrorMessage("Please only drop one file");
    }

    if (validateFile(files[0])) {
      setSelectedFile(files[0]);
      const reader = new FileReader();
      reader.readAsDataURL(files[0]);
      reader.onload = function (e) {
        setPreviewImage(e.target?.result as string);
      };
    } else {
      setErrorMessage("File is not of valid type");
    }
  };

  const dragOver = (e: React.DragEvent) => {
    e.preventDefault();
  };

  const dragEnter = (e: React.DragEvent) => {
    e.preventDefault();
  };

  const dragLeave = (e: React.DragEvent) => {
    e.preventDefault();
  };

  const fileDrop = (e: React.DragEvent) => {
    e.preventDefault();
    const files = e.dataTransfer.files;
    if (files.length) {
      handleFiles(files);
    }
  };

  const upload = () => {
    uploadFile({ variables: { description, image: selectedFile } });
  };

  console.log(loading, data);

  if (errorMessage) {
    return (
      <div>
        <p>{errorMessage}</p>
      </div>
    );
  }

  if (previewImage) {
    return (
      <div className="flex flex-col items-center w-1/3 p-4 mt-6 border-2">
        <img src={previewImage} />
        <textarea
          id="about"
          name="about"
          rows={3}
          onChange={(e) => setDescription(e.target.value)}
          className="relative block w-full px-3 py-2 mt-4 text-gray-900 placeholder-gray-500 border border-gray-300 appearance-none focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
        ></textarea>
        <button
          onClick={upload}
          className="relative flex justify-center w-full px-4 py-2 mt-4 font-bold text-white bg-blue-400 border border-transparent text-l group rounded-mdfocus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Upload
        </button>
      </div>
    );
  }

  return (
    <div
      className="flex justify-center w-1/3 p-4 mt-6 border-4 border-dashed border-light-blue-500"
      onDragOver={dragOver}
      onDragEnter={dragEnter}
      onDragLeave={dragLeave}
      onDrop={fileDrop}
    >
      <h2 className="text-gray-500">Drag to upload..</h2>
    </div>
  );
};

export default Upload;
