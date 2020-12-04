import Upload from "./Upload";
import { useQuery, gql } from "@apollo/client";
import { getPosts } from "./gql-types";

const GET_POSTS = gql`
  query getPosts($createdBefore: Timestamp) {
    getPosts(createdBefore: $createdBefore) {
      imageURL
      description
      author {
        firstName
        lastName
      }
    }
  }
`;

const Feed = () => {
  const { data } = useQuery<getPosts>(GET_POSTS);

  return (
    <div className="flex flex-col items-center">
      <div className="text-center">
        <h1 className="mt-6 text-4xl text-gray-700 font-dmsans">Kite</h1>
        <h2 className="text-gray-500">Influencing everyone, everywhere..</h2>
      </div>
      <div className="items-center justify-center w-1/3">
        <Upload />
      </div>
      <div className="w-1/3 mt-8">
        {data?.getPosts?.map((p) => (
          <div className="mt-3">
            <h3>{p?.author.firstName}</h3>
            <img src={p?.imageURL} />
            <p>{p?.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Feed;
