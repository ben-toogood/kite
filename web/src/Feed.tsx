import Upload from "./Upload";

const Feed = () => {
  return (
    <div>
      <div className="text-center">
        <h1 className="font-dmsans text-4xl text-gray-700 mt-6">Kite</h1>
        <h2 className="text-gray-500">Influencing everyone, everywhere..</h2>
      </div>
      <div className="flex flex-col justify-center w-full items-center">
        <Upload />
      </div>
    </div>
  );
};

export default Feed;
