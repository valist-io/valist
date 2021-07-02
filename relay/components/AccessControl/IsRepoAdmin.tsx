import { FunctionComponent } from 'react';
import useRepoAdmin from '../../hooks/useRepoAdmin';

const IsRepoAdmin:FunctionComponent<any> = (props) => {
  const isRepoAdmin = useRepoAdmin(props.orgName, props.repoName);
  if (isRepoAdmin) {
    return props.children;
  }
  return null;
};

export default IsRepoAdmin;
