# Dell Lookup
A (soon-to-be) web-interface for looking up warranty status based on Dell Service Tag written in Go.

### Configuration
Currently the only configuration option is to provide the application an API Key for Dell's Service Tag system. This file is stored at the root of the project.

config.json

```
{
    "ApiKey": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

### Dell Documentation
The official Dell Support Services API Documentation Can be found here:

http://en.community.dell.com/dell-groups/supportapisgroup/

### Legal
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

See the License for the specific language governing permissions and limitations under the License.
