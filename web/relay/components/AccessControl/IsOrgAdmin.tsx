import { FunctionComponent } from 'react';
import useOrgAdmin from '../../hooks/useOrgAdmin';

const IsOrgAdmin:FunctionComponent<any> = (props) => {
  const isOrgAdmin = useOrgAdmin(props.orgName);
  if (isOrgAdmin) {
    return props.children;
  }
  return null;
};

export default IsOrgAdmin;
