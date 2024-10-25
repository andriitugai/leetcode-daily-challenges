class Solution:
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        folder.sort()
        result = [folder[0]]

        for i in range(1, len(folder)):
            # Get the last added folder path and add a trailing slash
            last_folder = result[-1] + '/'
            
            # Check if current folder starts with last_folder
            # If it doesn't start with last_folder, then it's not a subfolder
            if not folder[i].startswith(last_folder):
                result.append(folder[i])
        
        return result